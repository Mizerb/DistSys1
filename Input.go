package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func (localN *Node) TweetEvent(message string) *Node {
	twt := tweet{message, localN.Id, localN.Id, time.Now().UTC(), localN.Ci, 2}

	//update the tweet in memory
	localN.Log[localN.Id] = append(localN.Log[localN.Id], twt)

	fmt.Println("Current messages in log:")
	for i := 0; i < len(localN.Log[localN.Id]); i++ {
		fmt.Println(" - ", localN.Log[localN.Id][i].Message)
	}
	fmt.Println("")

	//update the tweet in the physical log
	localN.writeLog()

	//Next TO DO:
	//send the log to the other ips

	return localN
}

func InputHandler(local *Node) {
	reader := bufio.NewReader(os.Stdin)
	for true {
		//done := make(chan bool)
		fmt.Printf("Please enter a Command: ")
		input, _ := reader.ReadString('\n')

		if i := strings.Index(input, "tweet"); i == 0 {
			message := input[6 : len(input)-1]
			fmt.Println("Tweet Called")
			local.TweetEvent(message)
		} else if i := strings.Index(input, "view"); i == 0 {
			fmt.Printf("View called\n")
		} else if i := strings.Index(input, "block"); i == 0 {
			username := input[6 : len(input)-1]
			fmt.Printf("Block called on %s\n", username)
			//create new tweet with type set to block, add to local dictonary
		} else if i := strings.Index(input, "unblock"); i == 0 {
			username := input[8 : len(input)-1]
			fmt.Printf("Unblock called on %s\n", username)
		} else if i := strings.Index(input, "exit"); i == 0 {
			fmt.Printf("Exit called, exiting...")
			break
		} else {
			fmt.Printf("Command not recognized\n")
		}
	}
}
