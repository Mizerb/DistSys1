package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const staticLog = "./localLog.json"
const staticDict = "./localDict.json"

type Node struct {
	id int
	Ci int

	userName string

	log      [][]tweet
	logMutex *sync.Mutex

	TimeArray [][]int
	TimeMutex *sync.Mutex

	blocks map[int]int //I guess I'm going to have to do an array of this, but I  don't want to
	// Let It Be Known To All Who Read the Comments, I Didn't Want This
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

	check, err := ret.LoadDict()
	if err != nil || check == false {
		//create static dict
		f, err := os.Create(staticDict)
		if err != nil {
			log.Fatal("Failed to create staticDict")
		}
		f.Close()
	}

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

func (n *Node) LoadDict() (bool, error) {
	_, err := os.Stat(staticDict)
	if os.IsNotExist(err) {
		log.Panicln("DICT FILE NOT YET CREATED")
		return false, nil
	}

	file, err := ioutil.ReadFile(staticDict)
	if err != nil {
		return false, err
	}
	n.blockMutex.Lock()
	defer n.blockMutex.Unlock()
	if err := json.Unmarshal(file, &n.blocks); err != nil {
		return false, err
	}
	return true, nil
}

func (n *Node) writeLog() {
	//n.logMutex.Lock()
	logBytes, err := json.Marshal(n.log)
	//defer n.logMutex.Unlock()
	if err != nil {
		log.Fatal(err)
		//might not want fatal here...
	}
	err = ioutil.WriteFile(staticLog, logBytes, 0644)
	if err != nil {
		log.Fatalln("Failed to write to staticlog")
	}
}

func (n *Node) writeDict() {
	//n.blockMutex.Lock()
	//defer n.blockMutex.Unlock()
	writeBytes, err := json.Marshal(n.blocks)
	if err != nil {
		log.Fatalf("failed to Marshel to static Dict\n")
	}
	err = ioutil.WriteFile(staticDict, writeBytes, 0644)
	if err != nil {
		//well shit,
		log.Fatalln("Failed to write to static dict")
		//
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
