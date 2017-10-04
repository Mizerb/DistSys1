package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TweetEvent(local *Node, message string) {
	fmt.Println("hiiiii")
	return
}

func InputHandler(local *Node) {
	reader := bufio.NewReader(os.Stdin)
	for true {
		//done := make(chan bool)
		fmt.Println("Please enter a Command:")
		input, _ := reader.ReadString('\n')

		if i := strings.Index(input, "tweet"); i == 0 {
			message := input[6 : len(input)-1]
			fmt.Println(message)
			//userTweet := tweet{message, myIP, time.Now().UTC()}
			//tweetUpdate(message, myIP)
			TweetEvent(local, message)
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
