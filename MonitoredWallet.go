package main

type MonitoredWallet struct {
	Monitored  Wallet `json:"monitored"`
	SinceBlock uint64 `json:"sinceBlock"`
}
