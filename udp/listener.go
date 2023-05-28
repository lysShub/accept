package udp

import (
	"net"
	"sync/atomic"
	"syscall"
)

type udpListener struct {
	laddr    *net.UDPAddr
	listener *net.UDPConn // unconnected udp conn

	b []byte

	closed atomic.Bool
}

var _ net.Listener = &udpListener{}

func NewListener(laddr *net.UDPAddr, mss int) (net.Listener, error) {
	if mss <= 0 {
		mss = 1536
	}
	var l = &udpListener{
		laddr: laddr,
		b:     make([]byte, mss),
	}

	var err error
	l.listener, err = net.ListenUDP("udp", laddr)
	if err != nil {
		return nil, err
	}

	rawConn, err := l.listener.SyscallConn()
	if err != nil {
		return nil, err
	}
	if err = setReuse(rawConn); err != nil {
		return nil, err
	}
	return l, nil
}

// Accept
//
//	will ignore null payload packet
func (l *udpListener) Accept() (net.Conn, error) {
	if l.closed.Load() {
		return nil, net.ErrClosed
	}

	n, raddr, err := l.listener.ReadFromUDP(l.b)
	if err != nil {
		return nil, err
	} else {

		var d = net.Dialer{
			LocalAddr: l.laddr,
			Control: func(network, address string, c syscall.RawConn) error {
				return setReuse(c)
			},
		}

		// BUG:
		//  before dial this connected conn, maybe recved more than one packet from raddr
		uconn, err := d.Dial("udp", raddr.String())
		if err != nil {
			return nil, err
		}

		var fp = make([]byte, n)
		copy(fp, l.b[:n])
		return &udpConn{Conn: uconn, firstPack: fp}, nil
	}
}

func (u udpListener) Close() error {
	if u.closed.CompareAndSwap(false, true) {
		defer func() {
			u.listener = nil
		}()
		return u.listener.Close()
	}
	return nil
}

func (u udpListener) Addr() net.Addr {
	return u.laddr
}
