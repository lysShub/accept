package main

import (
	"fmt"
	"net"

	"github.com/lysShub/accept/udp"
)

func main() {

	// 172.17.18.166 wsl
	// 172.31.1.244
	l, err := udp.NewListener(&net.UDPAddr{Port: 19986}, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go func(conn net.Conn) {
			var b = make([]byte, 1536)
			for {
				n, err := conn.Read(b)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("conned recv", string(b[:n]))
			}
		}(conn)
	}
}
