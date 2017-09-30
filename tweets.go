package main

import (
	"encoding/json"
	"log"
	"time"
)

type message struct {
	Ti     [][]int
	events [][]tweet
}

const ( //iota is reset to 0
	TWEET  = iota
	INSERT = iota // INSERT=0
	DELETE = iota // DELETE=1
)

type tweet struct {
	message  string
	user     int
	follower int
	clock    time.Time
	counter  int
}

func (n tweet) getTimestamp() time.Time {
	return n.clock
}

func (n tweet) getUser() int {
	return n.user
}

func (n tweet) getJSON() []byte {
	ret, err := json.Marshal(n)
	if err != nil {
		log.Printf("Failed to create JSON")
		return nil
	}
	return ret
}

func getTweets(msg []byte) ([]tweet, error) {
	var ret []tweet
	err := json.Unmarshal(msg, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
