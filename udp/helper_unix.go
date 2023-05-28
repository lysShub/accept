//go:build unix && linux
// +build unix,linux

package udp

import (
	"syscall"

	"golang.org/x/sys/unix"
)

func setReuse(conn syscall.RawConn) error {
	var e error
	err := conn.Control(func(fd uintptr) {
		e = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		if e != nil {
			return
		}
		e = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
		if e != nil {
			return
		}
	})
	if e != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
