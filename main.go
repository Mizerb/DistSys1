package main

import (
	"log"
	"os"
)

func main() {
	starting := os.Args[1]

	if _, err := os.Stat(starting); os.IsNotExist(err) {
		log.Fatalln("First arugment must configuration file")
	}

	local := makeNode(starting)
	InputHandler()
	return
}
