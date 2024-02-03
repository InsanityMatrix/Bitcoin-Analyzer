package main

type BitcoinAnalyzer struct {
	CurrentBlock     uint64  `json:"currentBlock"`
	MonitorThreshold float64 `json:"monitorThreshold"`
}

func (a *BitcoinAnalyzer) updateBlock(block uint64) {
	a.CurrentBlock = block
}
func (a *BitcoinAnalyzer) getCurrentBlock() uint64 {
	return a.CurrentBlock
}
