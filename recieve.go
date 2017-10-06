package main

/*
The logic behind this functiion starts with inserting and deleting things from
the blocks dict, but where does it go from there
The bit about removing things from the record is important
But I'll need to serach through events to figure out which one I need to remove it from
which will be a pain

Weird twists,
2 block commands
wait, shit what if a user block more than 1 person..
welcom to slice town, population me...
*/
func (n *Node) UpdateDict(events [][]tweet) {
	for _, stuff := range events {
		for _, record := range stuff {
			if record.Event == INSERT {
				//insert into dict I guess
				//yeah,

				//TODO handle user blocking mutlipul people
				// likley with arrays
				// for which I am sorry, but I have no choose
				// I wonder if someone's implemented that yet
				// would check but I don't have internet
				n.Blocks[record.User] = record.Follower
			} else if record.Event == DELETE {
				//handle VERY differently.
				// like I have to think with a a peice of paper
				// Because I can't just add it if it already exists

				delete(n.Blocks, record.User)
				//sure
			}
		}
	}
}

/*
UpdateLog put messages into log
wrote this because I don't want double nested loop in another function
*/
func (n *Node) UpdateLog(events [][]tweet) {
	//
	//n.logMutex.Lock()
	for i, noderec := range events {
		for _, record := range noderec {
			n.Log[i] = append(n.Log[i], record)
		}
	}
	//n.logMutex.Unlock()
	n.writeLog()
}

func (n *Node) receive(msg *message) {
	// Locks occur at this level
	// because if they occur any lower
	// one thread might change the clocks before another adds to the logs.
	// which would be ... very bad
	n.BlockMutex.Lock()
	defer n.BlockMutex.Unlock()

	n.LogMutex.Lock()
	defer n.LogMutex.Unlock()

	n.TimeMutex.Lock()
	defer n.TimeMutex.Unlock()
	//Figure which events are actually new
	newEvent := make([][]tweet, len(n.Log))
	for i := range n.Log {
		for j := range msg.Events[i] {
			if !(n.hasRec(msg.Events[i][j], n.Id)) {
				newEvent[i] = append(newEvent[i], msg.Events[i][j])
			}
		}
	}

	//update dictonary
	// crapppp

	n.UpdateDict(newEvent)

	//update the time array
	//n.TimeMutex.Lock()
	for k := range n.TimeArray[n.Id] {
		n.TimeArray[n.Id][k] = maxInt(n.TimeArray[n.Id][k], msg.Ti[msg.SendID][k])
	}
	for i := range n.TimeArray {
		for j := range n.TimeArray[i] {
			n.TimeArray[i][j] = maxInt(n.TimeArray[i][j], msg.Ti[i][j])
		}
	}
	//n.TimeMutex.Unlock()

	//update local log
	// by this point in time, all I should need to do is add the new events to the log
	// I should have verification at this point that they are infact, not currently in
	// the log
	n.UpdateLog(newEvent)

}
