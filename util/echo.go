package util

import (
	"log"
	"net"
)

/*
StartEchoTCPServer starts a TCP server that yells back at you.
*/
func StartEchoTCPServer() {
	conn, _ := net.Listen("tcp", "127.0.0.1:1984")

	log.Println("Echo TCP server listening at 127.0.0.1:1984")

	defer conn.Close()

	for {
		conn, _ := conn.Accept()
		defer conn.Close()

		go func(conn net.Conn) {
			for {
				buf := make([]byte, 65536)

				n, err := conn.Read(buf)
				if err != nil {
					return
				}

				conn.Write(buf[:n])
			}
		}(conn)
	}
}
