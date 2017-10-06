package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteLog(t *testing.T) {
	os.Remove(staticDict)
	os.Remove(staticLog)
	node := makeNode("entryData.json")

	msg := tweet{User: 0, Event: 0, Message: "haha"}
	node.log[0] = append(node.log[0], msg)
	node.writeLog()

	data, err := ioutil.ReadFile(staticLog)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	var check [][]tweet
	err = json.Unmarshal(data, check)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(check) != len(node.log) {
		fmt.Println("log and file not same")
	}
	for i := range node.log {
		for j := range node.log[i] {
			if node.log[i][j].Message != check[i][j].Message {
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
