package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type server struct {
	n    int
	ip   string
	port string
}

func echoSender(msg string, meh server, ch chan<- bool) {
	// do the thing
	conn, err := net.Dial("tcp", meh.ip+meh.port)
	if err != nil {
		//oh shit
		fmt.Printf("Failed to connect to %d\n", meh.n)
		ch <- false
	}
	defer conn.Close()
	fmt.Fprintf(conn, msg)
	//status, err := bufio.NewReader(conn).ReadString('\n')
	var echo string
	//readme := bufio.NewReader(conn)
	//echo, _ = readme.ReadString('\n')

	_, err = fmt.Fscanln(conn, echo)
	//_, err = fmt.Fscanf(conn, "%s", &echo)
	if err != nil {
		//oh shit
		//fmt.Printf("%d  %d \n", len(echo), len(msg))
		fmt.Printf("Execpected \"%s\", recieved \"%s\"\n", msg, echo)
		ch <- false
	}
	fmt.Printf("Message from %d: %s\n", meh.n, echo)

	ch <- true
}

func recieveSendEcho(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 200)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal("Failed to recieve msg")
	}

	m, err := conn.Write(buffer)
	if n != m && err != nil {
		log.Fatal("Failed to Send Message")
	}
	log.Printf("Successfully echoed Message %s", buffer)
	return
}

func echoReciever(meh server) {
	ln, err := net.Listen("tcp", meh.port)
	if err != nil {
		//crap
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			//crap
			log.Fatal(err)
		}
		go recieveSendEcho(conn)
	}
}

func GetPosition(fileName string, n int) []server {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	servers := make([]server, 3)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		var serv server
		serv.n = i
		fmt.Sscanf(scanner.Text(), "%s %s", &serv.ip, &serv.port)
		servers[i] = serv
		i++
	}
	fmt.Println(servers)
	return servers
}

/*
func Read_File(file_name string) string int{
	return nil, nil
}
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage:\n 1: Interger to designate sever number\n 2:File location for hosts.txt\n")
		return
	}
	arg := os.Args[1]
	fmt.Println(arg)
	n, err := strconv.Atoi(arg)
	_, check := os.Stat(os.Args[2])
	if os.IsNotExist(check) && err != nil {
		fmt.Print("Usage:\n 1: Interger to designate sever number\n 2:File location for hosts.txt\n")
		return
	}

	servers := GetPosition(os.Args[2], n)
	fmt.Println(servers)
	if n == 0 {
		reader := bufio.NewReader(os.Stdin)
		for /*i := 0; i < 2; i++*/ {
			done := make(chan bool)
			fmt.Printf("%c enter test: ", '%')
			input, _ := reader.ReadString('\n')
			//fmt.Println(text)

			go echoSender(input, servers[1], done)
			go echoSender(input, servers[2], done)

			if <-done && <-done {
				//nothing, for now
			}
		}
	} else {
		echoReciever(servers[n])
	}
}
