package main

import (
	"encoding/json"
	"log"
	"net"
)

//will generate messages that can be sent else were
// but main purpose is to generate messages for all
// other nodes
// essentailly, you only can send out messages when
//  there's a new tweet, so this requires that tweet.
// I'll talk it over with Ian...
func (n *Node) BroadCast() {
	// locks should be applied here
	/*
		n.BlockMutex.Lock()
		defer n.BlockMutex.Unlock()

		n.LogMutex.Lock()
		defer n.LogMutex.Unlock()

		n.TimeMutex.Lock()
		defer n.TimeMutex.Unlock()
	*/
	for i, ip := range n.IPtargets {
		if ok := n.Blocks[n.Id][i]; ok {
			log.Println("ID ", i, " is blocked, not sending to location")
			continue
		}
		conn, err := net.Dial("tcp", ip)
		if err != nil {
			log.Println("Failed to connect to ", ip, "  ", err)
			continue
		}
		n.Send(conn, i)
	}
	return
}

func (n *Node) Send(conn net.Conn, k int) {
	defer conn.Close()
	var msg message
	//n.LogMutex.Lock()
	msg.Events = make([][]tweet, len(n.Log))
	msg.Ti = n.TimeArray
	msg.SendID = n.Id

	for i := range n.Log {
		for j := range n.Log[i] {
			if !n.hasRec(n.Log[i][j], k) {
				msg.Events[i] = append(msg.Events[i], n.Log[i][j])
			}
		}
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Println("failed to build message for ", k, "   ", err)
		return
	}

	check, err := conn.Write(bytes)
	if err != nil || check != len(bytes) {
		log.Println("Failed to send message to ", k, "  ", err)
		return
	}
	log.Println("Successfully sent message to ", k)
}
