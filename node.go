package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const staticLog = "./localLog.json"

type Node struct {
	id int
	Ci int

	log      [][]tweet
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

	ret.log = make([][]tweet, info.totalNodes)
	for i := 0; i < info.totalNodes; i++ {
		ret.log[i] = make([]tweet, 0, 10)
	}
	ret.logMutex = &sync.Mutex{}

	if _, err := ret.LoadTweets(staticLog); err != nil {
		log.Fatal("welp")
	}

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

func (n *Node) LoadTweets(filename string) (bool, error) {
	_, err := os.Stat(staticLog)
	if os.IsNotExist(err) {
		log.Println("LOG FILE NOT YET CREATED")
		return false, nil
	}

	file, err := ioutil.ReadFile(staticLog)
	if err != nil {
		return false, err
	}

	n.logMutex.Lock()
	defer n.logMutex.Unlock()
	if err := json.Unmarshal(file, &n.log); err != nil {
		return false, err
	}

	return true, nil
}

func (n *Node) hasRec(msg tweet, k int) bool {
	n.TimeMutex.Lock()
	ret := n.TimeArray[k][msg.user] >= msg.counter
	n.TimeMutex.Unlock()
	return ret
}

//will generate messages that can be sent else were
// but main purpose is to generate messages for all
// other nodes
// essentailly, you only can send out messages when
//  there's a new tweet, so this requires that tweet.
// I'll talk it over with Ian...
func (n *Node) BroadCast(msg tweet) {
	return
}
