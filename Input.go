package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
type Rankings struct {
	Keyword  string `json:"keyword"`
	GetCount uint32 `json:"get_count"`
	Engine   string `json:"engine"`
	Locale   string `json:"locale"`
	Mobile   bool   `json:"mobile"`
}
*/

func (localN *Node) TweetEvent(message string) *Node {
	//fmt.Println("hiiiii")
	//get the id of
	/*file, err := ioutil.ReadFile("./entryData.json")
	if err != nil {
		return
	}
	if err := json.Unmarshal(file, &localN.log); err != nil {
		return
	}*/
	twt := tweet{message, localN.Id, localN.Id, time.Now().UTC(), localN.Ci, 2}

	fmt.Println("Current messages in log:")
	for i := 0; i < len(localN.Log[localN.Id]); i++ {
		fmt.Println(" - ", localN.Log[localN.Id][i].Message)
	}
	fmt.Println("")

	//update the tweet in memory
	localN.Log[localN.Id] = append(localN.Log[localN.Id], twt)

	localN.writeLog()
	//update the tweet in the physical log
	//THIS IS A TEST

	/*
		var jsonBlob = []byte(`{"keyword":"hipaa compliance form", "get_count":157, "engine":"google", "locale":"en-us", "mobile":false}`)
		rankings := Rankings{}
		err := json.Unmarshal(jsonBlob, &rankings)
		if err != nil {
			panic(err)
		}
		rankingsJson, _ := json.Marshal(rankings)
		err = ioutil.WriteFile(staticLog, rankingsJson, 0644)
		fmt.Printf("%+v", rankings)
	*/

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
			//userTweet := tweet{message, myIP, time.Now().UTC()}
			//tweetUpdate(message, myIP)
			//local := TweetEvent(local, message)
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
