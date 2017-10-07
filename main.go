package main

import (
	"fmt"
	"log"
	"os"
)

/*
Next To Do:
	- Implement wuu-berstien for log consolidation
	- Many other things...
	- Test sending & recieving
	- Truncate as required for added and deleted dict entries
	- Ensure logic with muteses is correct (things aren't overwritten at bad times)
	- find other things to do....
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
	go listen(local)
	InputHandler(local)
	return
}
