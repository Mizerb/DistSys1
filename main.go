package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type node struct {
}

const ( //iota is reset to 0
	INSERT = iota // INSERT=0
	DELETE = iota // DELETE=1
)



func main() {

	reader := bufio.NewReader(os.Stdin)
	for true {
		//done := make(chan bool)
		fmt.Println("Please enter a Command:")
		input, _ := reader.ReadString('\n')
		if i := strings.Index(input, "tweet"); i < 1 {
			fmt.Printf("Found at %d , TWEET\n", i)
		} else if i := strings.Index(input, "view"); i < 1 {
			fmt.Printf("great, now you get to see my history")
		} else if i := strings.Index(input, "block"); i < 1 {
			fmt.Printf("HOW DIRE!\n")
		} else if i := strings.Index(input, "unblock"); i < 1 {
			fmt.Printf("HOW FRIENDLY\n")
		} else {
			fmt.Printf("Command not recognized\n")
		}
	}
	return
}
