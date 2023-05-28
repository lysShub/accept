//go:build windows
// +build windows

package udp

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func setReuse(conn syscall.RawConn) error {
	var e error
	err := conn.Control(func(fd uintptr) {
		e = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
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
