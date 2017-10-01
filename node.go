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

	userName string

	log      [][]tweet
	logMutex *sync.Mutex

	TimeArray [][]int
	TimeMutex *sync.Mutex

	blocks     map[int]int
	blockMutex *sync.Mutex

	listenPort int
	IPtargets  map[int]string
}

func makeNode(inputfile string) *Node {
	ret := new(Node)

	file, err := ioutil.ReadFile(inputfile)
	if err != nil {
		log.Fatal("Cannot Open file " + inputfile)
	}

	type startinfo struct {
		id         int
		localport  int
		totalNodes int
		IPs        map[string]string
	}

	var info startinfo //Deserialize the JSON
	if err := json.Unmarshal(file, &info); err != nil {
		log.Fatal(err)
	}

	ret.listenPort = info.localport
	ret.id = info.id

	ret.log = make([][]tweet, info.totalNodes)
	for i := 0; i < info.totalNodes; i++ {
		ret.log[i] = make([]tweet, 0, 10)
	}
	ret.logMutex = &sync.Mutex{}

	if check, err := ret.LoadTweets(staticLog); err != nil || check == false {
		//log.Fatal("welp")
		//create file
		f, err := os.Create(staticLog)
		if err != nil {
			log.Fatal("cannot create log")
		}
		f.Close()
	}

	ret.TimeArray = make([][]int, info.totalNodes)
	/*for i := range ret.TimeArray {
		ret.TimeArray[i] = make([]int, info.totalNodes)
	}*/
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

func (n *Node) writeLog() {
	n.logMutex.Lock()
	logBytes, err := json.Marshal(n.log)
	defer n.logMutex.Unlock()
	if err != nil {
		log.Fatal(err)
		//might not want fatal here...
	}
	err = ioutil.WriteFile(staticLog, logBytes, 0644)
	if err != nil {
		log.Fatalln("Failed to write to staticlog")
	}

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

func (n *Node) receive(msg *message) {
	//Figure which events are actually new
	newEvent := make([][]tweet, len(n.log))
	for i := range n.log {
		for j := range msg.events[i] {
			if !(n.hasRec(msg.events[i][j], n.id)) {
				newEvent[i] = append(newEvent[i], msg.events[i][j])
			}
		}
	}

	//update dictonary
	//

	//update the time array
	n.TimeMutex.Lock()
	for k := range n.TimeArray[n.id] {
		n.TimeArray[n.id][k] = maxInt(n.TimeArray[n.id][k], msg.Ti[msg.sendID][k])
	}
	for i := range n.TimeArray {
		for j := range n.TimeArray[i] {
			n.TimeArray[i][j] = maxInt(n.TimeArray[i][j], msg.Ti[i][j])
		}
	}
	n.TimeMutex.Unlock()

	//update local log

}
