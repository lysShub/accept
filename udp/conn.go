package udp

import (
	"net"
	"sync"
	"syscall"
)

type udpConn struct {
	sync.Once
	net.Conn

	firstPack []byte
}

func (u *udpConn) Read(b []byte) (n int, err error) {
	u.Do(func() {
		n = copy(b, u.firstPack)
		if n < len(u.firstPack) {
			err = syscall.EMSGSIZE
		}
		u.firstPack = nil
	})

	if n == 0 {
		return u.Conn.Read(b)
	} else {
		return n, err
	}
}
