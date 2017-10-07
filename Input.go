package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

//check to see if tweet, block or unblock event

func reverse(logArray []tweet) []tweet {
	for i, j := 0, len(logArray)-1; i < j; i, j = i+1, j-1 {
		logArray[i], logArray[j] = logArray[j], logArray[i]
	}
	return logArray
}

func (localN *Node) ViewTweets() {
	fmt.Println("Current messages in log:")
	logReverse := reverse(localN.Log[localN.Id])
	for i := 0; i < len(logReverse); i++ {
		if logReverse[i].Event == 0 {
			//TO DO: Check the dictionary to see if the user is currently blocked
			fmt.Println(" - ", logReverse[i].Message)
		}
	}
}

func (localN *Node) TweetEvent(message string) {
	twt := tweet{message, localN.Id, localN.Id, time.Now().UTC(), localN.Ci, 0}

	//update the tweet in memory
	localN.Log[localN.Id] = append(localN.Log[localN.Id], twt)

	//update the tweet in the physical log
	localN.writeLog()

	//Update the counter
	localN.Ci++

	//send the log to the other ips
	localN.BroadCast()
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
			local.ViewTweets()
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
