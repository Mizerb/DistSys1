package main

import (
	"fmt"
	"log"
	"os"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("First argument must be configuration file")
	}
	starting := os.Args[1]

	if _, err := os.Stat(starting); os.IsNotExist(err) {
		log.Fatalln("Unable to open configuration file\nPlease confirm it's a json")
	}

	local := makeNode(starting)
	//go listen(local)
	InputHandler(local)
	return
}
