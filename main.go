package main

import "time"

const ( //iota is reset to 0
	INSERT = iota // INSERT=0
	DELETE = iota // DELETE=1
)

type tweet struct {
	message string
	user    int
	clock   time.Time
}

func main() {
	return
}
