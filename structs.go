package main

import (
	"encoding/json"
	"log"
	"time"
)

type eR interface {
	getTimestamp() time.Time
	getUser() int
	getJSON() []byte
}

type tweet struct {
	message string
	user    int
	clock   time.Time
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

type block struct {
	user     int
	follower int
	clock    time.Time
}

func (n block) getTimestamp() time.Time {
	return n.clock
}

func (n block) getUser() int {
	return n.user
}

func (n block) getJSON() []byte {
	ret, err := json.Marshal(n)
	if err != nil {
		log.Printf("Failed to create JSON")
		return nil
	}
	return ret
}

type unblock struct {
	targetuser int
	user       int
	clock      time.Time
}

func (n unblock) getTimestamp() time.Time {
	return n.clock
}

func (n unblock) getUser() int {
	return n.user
}

func (n unblock) getJSON() []byte {
	ret, err := json.Marshal(n)
	if err != nil {
		log.Printf("Failed to create JSON")
		return nil
	}
	return ret
}
