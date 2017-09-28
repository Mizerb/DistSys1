package main

import (
	"bufio"
	"fmt"
	"log"
	//"log"
	"os"
	"strings"
	//"net"
	//time
)

/*
Next To Do:
- Figure out how to log the information. Issue related to format and how it's going to get passed between locations
  - Might be a good idea to store in different file format, rather than golang's log functionality
- Distribute the log passing between different locations (after tweet is called)
- Implement wuu-berstien for log consolidation
- Set up dictionary for storing of block and unblock commands
- Many other things...
*/

type node struct {
}

const ( //iota is reset to 0
	TWEET  = iota
	INSERT = iota // INSERT=0
	DELETE = iota // DELETE=1
)

func tweetUpdate(message string, myIP string) {
	log.Printf("%s,%s,%s", "tweet", myIP, message)
	//log.Printf("%s   %s   %s", userTweet.user, userTweet.message, userTweet.clock)
	return
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for true {

		//Basic code for calling to recieve tweets (and logs) from specified locations
		/*
			siteNum := 2
			var sites [2]string
			sites[0] = "127.0.0.1:7777"
			sites[1] = "127.0.0.1:7777"

			for i := 0; i < siteNum; i++ {
				//go listen(sites[i])
			}
		*/
		myIP := "127.0.0.1"
		//var myPort string = "7777"

		//done := make(chan bool)
		fmt.Println("Please enter a Command:")
		input, _ := reader.ReadString('\n')

		if i := strings.Index(input, "tweet"); i == 0 {
			message := input[6 : len(input)-1]
			fmt.Println(message)
			//userTweet := tweet{message, myIP, time.Now().UTC()}
			tweetUpdate(message, myIP)
		} else if i := strings.Index(input, "view"); i == 0 {
			fmt.Printf("View called\n")
		} else if i := strings.Index(input, "block"); i == 0 {
			username := input[6 : len(input)-1]
			fmt.Printf("Block called on %s\n", username)
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
	return
}
