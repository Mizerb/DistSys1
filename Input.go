package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func OrganizeTweets(logContent [][]tweet) []tweet {
	var combineList []tweet
	for i := 0; i < len(logContent); i++ {
		for j := 0; j < len(logContent[i]); j++ {
			if len(combineList) == 0 {
				combineList = append(combineList, logContent[i][j])
				continue
			} else {
				for k := 0; k < len(combineList); k++ {
					if combineList[k].Clock.After(logContent[i][j].Clock) {
						combineListBefore := combineList[0:k]
						combineListAfter := combineList[k : len(combineList)-1]
						combineList = append(combineListBefore, logContent[i][j])
						for l := 1; l < len(combineListAfter); l++ {
							combineList = append(combineList, combineListAfter[l])
						}
						break
					} else if k+1 == len(combineList) {
						combineList = append(combineList, logContent[i][j])
						break
					}
				}
			}
		}
	}
	return combineList
}

func (localN *Node) ViewTweets() {
	fmt.Println("Current events in log:")
	organizedLog := OrganizeTweets(localN.Log)
	logReverse := reverse(organizedLog)
	for i := 0; i < len(logReverse); i++ {
		//TO DO: Check the dictionary to see if the user is currently blocked
		fmt.Printf(time.Time.String(logReverse[i].Clock) + " - ")
		fmt.Printf("User " + strconv.Itoa(logReverse[i].User) + " at counter " + strconv.Itoa(logReverse[i].Counter) + ", ")
		if logReverse[i].Event == 0 {
			fmt.Printf("TWEET: " + logReverse[i].Message)
		} else if logReverse[i].Event == 1 {
			fmt.Printf("BLOCK: Follower " + strconv.Itoa(logReverse[i].Follower))
		} else if logReverse[i].Event == 2 {
			fmt.Printf("UNBLOCK: Follower " + strconv.Itoa(logReverse[i].Follower))
		}
		//fmt.Println(" - ", logReverse[i].Message)
		fmt.Println("")
	}
}

func (localN *Node) TweetEvent(message string) {
	//Update the counter
	localN.Ci++
	twt := tweet{message, localN.Id, localN.Id, time.Now().UTC(), localN.Ci, 0}
	//update the tweet in memory
	localN.Log[localN.Id] = append(localN.Log[localN.Id], twt)
	//update the tweet in the physical log
	localN.writeLog()
	//send the log to the other ips
	localN.BroadCast()
}

func (localN *Node) InvalidBlock(username string, blockType int) bool {
	//SAFTEY CHECKS TO IMPLEMENT:
	// - User calls block on another user that exists
	// - User calls block on a user that is already blocked
	// - User doesn't call block on themself
	// - User calls unblock on another user that exists
	// - User calls unblock on a user that is not in the dictionary
	return false
}

func (localN *Node) BlockUser(username string) {
	if localN.InvalidBlock(username, 1) == true {
		log.Println("Invalid Block Call")
		return
	}
	localN.Ci++
	userID, _ := strconv.Atoi(username)
	twtBlock := tweet{"", localN.Id, userID, time.Now().UTC(), localN.Ci, 1}
	localN.Log[localN.Id] = append(localN.Log[localN.Id], twtBlock)
	localN.Blocks[localN.Id][userID] = true
	localN.writeLog()
	localN.writeDict()
}

func (localN *Node) UnblockUser(username string) {
	if localN.InvalidBlock(username, 2) == true {
		log.Println("Invalid Block Call")
		return
	}
	userID, _ := strconv.Atoi(username)
	twtUnblock := tweet{"", localN.Id, userID, time.Now().UTC(), localN.Ci, 2}
	localN.Log[localN.Id] = append(localN.Log[localN.Id], twtUnblock)
	delete(localN.Blocks[localN.Id], userID)
	localN.writeLog()
	localN.writeDict()
	localN.Ci++
}

func InputHandler(local *Node) {
	reader := bufio.NewReader(os.Stdin)
	for true {
		//done := make(chan bool)
		fmt.Printf("Please enter a Command: ")
		inputTmp, _ := reader.ReadString('\n')
		input := strings.Replace(inputTmp, "\r", "", -1)

		if i := strings.Index(input, "tweet"); i == 0 {
			message := input[6 : len(input)-1]
			fmt.Println("Tweet Called")
			local.TweetEvent(message)
		} else if i := strings.Index(input, "view"); i == 0 {
			fmt.Printf("View called\n")
			local.ViewTweets()
		} else if i := strings.Index(input, "block"); i == 0 {
			//username := input[6 : len(input)-1]
			username := input[6:7]
			fmt.Printf("Block called on %s\n", username)
			local.BlockUser(username)
		} else if i := strings.Index(input, "unblock"); i == 0 {
			//username := input[8 : len(input)-1]
			username := input[8:9]
			fmt.Printf("Unblock called on %s\n", username)
			local.UnblockUser(username)
		} else if i := strings.Index(input, "exit"); i == 0 {
			fmt.Printf("Exit called, exiting...")
			break
		} else {
			fmt.Printf("Command not recognized\n")
		}
	}
}
