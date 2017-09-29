package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Node struct {
	id int
	Ci int

	log      []tweet
	logMutex *sync.Mutex

	TimeArray [][]int
	TimeMutex *sync.Mutex

	blocks     map[int]int
	blockMutex *sync.Mutex

	IPtargets map[int]string
}

func makeNode(inputfile string) *Node {
	ret := new(Node)

	file, err := ioutil.ReadFile(inputfile)
	if err != nil {
		log.Fatal("Cannot Open file " + inputfile)
	}

	type startinfo struct {
		id         int
		totalNodes int
		IPs        []string
	}

	var info startinfo
	if err := json.Unmarshal(file, &info); err != nil {
		panic(err)
	}

	ret.id = info.id

	ret.log = make([]tweet, 1)
	ret.logMutex = &sync.Mutex{}

	ret.TimeArray = make([][]int, info.totalNodes)
	for i := range ret.TimeArray {
		ret.TimeArray[i] = make([]int, info.totalNodes)
	}
	ret.TimeMutex = &sync.Mutex{}

	ret.blocks = make(map[int]int)
	ret.blockMutex = &sync.Mutex{}

	ret.IPtargets = make(map[int]string)

	ret.Ci = 0

	return ret
}

func (n *Node) LoadTweets(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Println("LOG FILE NOT YET CREATED")
		return nil
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (n *Node) hasRec(msg tweet, k int) bool {
	n.TimeMutex.Lock()
	ret := n.TimeArray[k][msg.user] >= msg.counter
	n.TimeMutex.Unlock()
	return ret
}
