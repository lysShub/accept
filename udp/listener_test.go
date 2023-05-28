package udp

import (
	"fmt"
	"net"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBase(t *testing.T) {

	l, err := NewListener(&net.UDPAddr{IP: net.ParseIP("172.17.18.166"), Port: 19986}, 0)
	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			t.Fatal(err)
		}

		go func(conn net.Conn) {
			var b = make([]byte, 1536)
			for {
				n, err := conn.Read(b)
				if err != nil {
					t.Fatal(err)
				}
				fmt.Println("conned recv", string(b[:n]))
			}
		}(conn)
	}
}

func TestBoundary(t *testing.T) {
	var saddr = &net.UDPAddr{Port: 19986}
	var conns atomic.Int32
	var packs atomic.Int32

	// server
	go func() {
		l, err := NewListener(saddr, 0)
		require.NoError(t, err)

		for {
			conn, err := l.Accept()
			require.NoError(t, err)
			conns.Add(1)

			go func(conn net.Conn) {
				var b = make([]byte, 1536)
				for {
					_, err := conn.Read(b)
					require.NoError(t, err)
					packs.Add(1)
				}
			}(conn)
		}
	}()

	conn, err := net.DialUDP("udp", nil, saddr)
	require.NoError(t, err)
	for i := 0; i < 1e4; i++ {
		_, err = conn.Write([]byte{byte(i)})
		require.NoError(t, err)
	}

	require.Equal(t, 1, conns.Load())
}
