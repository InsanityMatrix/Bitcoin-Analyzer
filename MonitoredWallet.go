package main

type MonitoredWallet struct {
	Monitored  Wallet `json:"monitored"`
	SinceBlock uint64 `json:"sinceBlock"`
}

func (m *MonitoredWallet) isDead(blockNum uint64) bool {
	//Get the last transaction block and see if it has been 2 week since the last transaction
	var blocksInTwoWeeks uint64 = 2 * 7 * 24 * 60 / 10
	if m.Monitored.Transactions[len(m.Monitored.Transactions)-1].TransactionBlock+blocksInTwoWeeks < blockNum {
		return true
	}
	return false
}
func (m *MonitoredWallet) isEval(blockNum uint64) bool {
	//Check if it has been monitored for atleast 3 days - in blocks
	var blocksInThreeDays uint64 = 3 * 24 * 60 / 10
	if m.SinceBlock+blocksInThreeDays < blockNum {
		return true
	}
	return false
}

// TODO: A more refined exchange evaluation algorithm
func (m *MonitoredWallet) evalExchange(blockNum uint64) bool {
	//Check if wallet is exchange
	if m.Monitored.IsExchange {
		return true
	}
	//For starters, if it has more than 20 transactions per day since monitored, it is an exchange
	var blocksInOneDay uint64 = 24 * 60 / 10
	totalTransactions := uint64(len(m.Monitored.Transactions))
	blocksPassed := blockNum - m.SinceBlock
	daysPassed := blocksPassed / blocksInOneDay
	if totalTransactions/daysPassed > 20 {
		m.Monitored.IsExchange = true
		return true
	}
	//If not, check if it has more than 120 transactions in total
	if totalTransactions > 120 {
		m.Monitored.IsExchange = true
		return true
	}
	return false
}
