package main

import "net"

func listen(port string, serv *Node) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		//crappp
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			//crappppp
		}
		go handleConn(conn, serv)
	}
}

func handleConn(conn net.Conn, serv *Node) {

	return
}
