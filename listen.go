package main

import "net"

func listen(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		//crappp
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			//crappppp
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	return
}
