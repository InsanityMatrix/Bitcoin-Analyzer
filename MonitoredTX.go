package main

type MonitoredTX struct {
	StartAddress string        `json:"startAddress"`
	EndAddress   string        `json:"endAddress"`
	Transactions []Transaction `json:"transactions"`
	LastBlock    uint64        `json:"startBlock"` //The last block a transaction was added to this chain
}

func (m *MonitoredTX) addTransaction(transaction Transaction) {
	m.Transactions = append(m.Transactions, transaction)
}
func (m *MonitoredTX) isDead(CurrentBlock uint64) bool {
	if m.LastBlock+2016 < CurrentBlock {
		return true
	}
	return false
}
