package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type cheat struct {
	Message string
	//user     int
	//follower int
	//clock    time.Time
	//counter int
	//event   int
}

func TestWriteLog(t *testing.T) {
	os.Remove(staticDict)
	os.Remove(staticLog)
	node := makeNode("entryData.json", 0)

	fmt.Println("Listening on ", node.ListenPort)

	msg := tweet{User: 0, Event: 0, Message: "haha"}
	node.Log[0] = append(node.Log[0], msg)
	node.writeLog()

	data, err := ioutil.ReadFile(staticLog)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	var check [][]tweet
	err = json.Unmarshal(data, &check)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(check) != len(node.Log) {
		fmt.Println("log and file not same")
		t.Fail()
	}

	if len(check[0]) != len(node.Log[0]) {
		fmt.Println("log and file not same")
		t.Fail()
	}
	for i := range node.Log {
		for j := range node.Log[i] {
			if node.Log[i][j].Message != check[i][j].Message {
				fmt.Println("log and file dont' contain same data")
				t.Fail()
			}
		}
	}
}

func TestMarsheling(t *testing.T) {
	os.Remove(staticDict)
	os.Remove(staticLog)
	//node := makeNode("entryData.json")

	msg := tweet{Message: "crap"}
	data, err := json.Marshal(msg)
	fmt.Println(string(data))
	fmt.Println(msg)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	var check cheat
	err = json.Unmarshal(data, &check)
	if check.Message != msg.Message {
		fmt.Println("marshel did not carry data")
		t.Fail()
	}

	//node.log[0] = append(node.log[0], msg)

}

func TestTruncate(t *testing.T) {
	os.Remove(staticDict)
	os.Remove(staticLog)

	local := makeNode("testentry.json", 0)
	local.incrementClock()
	twtBlock := tweet{"", local.Id, 0, time.Now().UTC(), local.Ci, 1}
	local.Log[local.Id] = append(local.Log[local.Id], twtBlock)
	local.Blocks[local.Id][0] = true

	local.incrementClock()
	twtUnblock := tweet{"", local.Id, 0, time.Now().UTC(), local.Ci, 2}
	local.Log[local.Id] = append(local.Log[local.Id], twtUnblock)
	local.Blocks[local.Id][0] = false

	//local.incrementClock()
	//taw := tweet{"bob", local.Id, 99, time.Now().UTC(), local.Ci, 0}
	//local.Log[local.Id] = append(local.Log[local.Id], taw)
	local.writeDict()
	local.CleanDict()
	local.writeDict()
	if len(local.Log[local.Id]) != 0 {
		fmt.Println("Log not equal to 0")
		fmt.Println(local.Log)
		t.Fail()
	}

}
