package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	if len(os.Args) < 3 {
		fmt.Println("First argument must be configuration file, Second must be the ID of the location")
	}
	starting := os.Args[1]
	id := os.Args[2]

	if _, err := os.Stat(starting); os.IsNotExist(err) {
		log.Fatalln("Unable to open configuration file\nPlease confirm it's a json")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln("Integer not entered for second argument (id)")
	}

	local := makeNode(starting, idInt)
	go listen(local)
	InputHandler(local)
	return
}
