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
	Id int
	Ci int

	UserName string

	Log      [][]tweet
	LogMutex *sync.Mutex

	TimeArray [][]int
	TimeMutex *sync.Mutex

	Blocks map[int]map[int]bool //I guess I'm going to have to do an array of this, but I  don't want to
	// Let It Be Known To All Who Read the Comments, I Didn't Want This
	BlockMutex *sync.Mutex

	ListenPort int
	IPtargets  map[int]string
}

func makeNode(inputfile string) *Node {
	ret := new(Node)

	file, err := ioutil.ReadFile(inputfile)
	if err != nil {
		log.Fatal("Cannot Open file " + inputfile)
	}

	type startinfo struct {
		Id         int
		Localport  int
		TotalNodes int
		IPs        map[string]string
	}

	var info startinfo //Deserialize the JSON
	if err := json.Unmarshal(file, &info); err != nil {
		log.Fatal(err)
	}

	//To get the id, localport, and totalNodes, this separate object needs to be used
	var dat map[string]interface{}
	if err := json.Unmarshal(file, &dat); err != nil {
		panic(err)
	}

	info.Id = int(dat["Id"].(float64))
	info.Localport = int(dat["Localport"].(float64))
	info.TotalNodes = int(dat["TotalNodes"].(float64))

	ret.ListenPort = info.Localport
	ret.Id = info.Id

	ret.Log = make([][]tweet, info.TotalNodes)
	for i := 0; i < info.TotalNodes; i++ {
		ret.Log[i] = make([]tweet, 0, 10)
	}
	ret.LogMutex = &sync.Mutex{}

	if check, err := ret.LoadTweets(staticLog); err != nil || check == false {
		//log.Fatal("welp")
		//create file
		f, err := os.Create(staticLog)
		if err != nil {
			log.Fatal("cannot create log")
		}
		f.Close()
	}

	ret.TimeArray = make([][]int, info.TotalNodes)
	/*for i := range ret.TimeArray {
		ret.TimeArray[i] = make([]int, info.totalNodes)
	}*/
	ret.TimeMutex = &sync.Mutex{}

	ret.Blocks = make(map[int]map[int]bool)
	ret.BlockMutex = &sync.Mutex{}
	for i := 0; i < info.TotalNodes; i++ {
		ret.Blocks[i] = make(map[int]bool)
	}

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

	n.LogMutex.Lock()
	defer n.LogMutex.Unlock()
	if err := json.Unmarshal(file, &n.Log); err != nil {
		return false, err
	}

	return true, nil
}

func (n *Node) LoadDict() (bool, error) {
	_, err := os.Stat(staticDict)
	if os.IsNotExist(err) {
		//log.Panicln("DICT FILE NOT YET CREATED")
		log.Println("DICT FILE NOT YET CREATED")
		return false, nil
	}

	file, err := ioutil.ReadFile(staticDict)
	if err != nil {
		return false, err
	}
	n.BlockMutex.Lock()
	defer n.BlockMutex.Unlock()
	if err := json.Unmarshal(file, &n.Blocks); err != nil {
		return false, err
	}
	return true, nil
}

func (n *Node) writeLog() {
	//n.logMutex.Lock()
	logBytes, err := json.Marshal(n.Log)
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
	writeBytes, err := json.Marshal(n.Blocks)
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
	ret := n.TimeArray[k][msg.User] >= msg.Counter
	n.TimeMutex.Unlock()
	return ret
}
