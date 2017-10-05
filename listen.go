package main

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
)

func listen(serv *Node) {
	ln, err := net.Listen("tcp", strconv.Itoa(serv.listenPort))
	defer ln.Close()
	if err != nil {
		//crappp
		log.Fatalln("Failed to connect on port, shutting down")
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
	defer conn.Close()
	var data []byte
	_, err := conn.Read(data)
	if err != nil {
		log.Println(err)
		return
	}
	var msg *message
	err = json.Unmarshal(data, msg)
	if err != nil {
		log.Println(err)
		return
	}

	//serv.receive(msg)
	return
}
