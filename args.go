package main

import (
	"flag"
)

func getArgs() (string, uint64) {
	//Define Flags

	datadirPtr := flag.String("datadir", "data", "Path to the data directory")
	startblock := flag.Uint64("startblock", uint64(0), "Block to start indexing from - Default starts indexing from latest block in datadir or on blockchain")

	flag.Parse()
	return *datadirPtr, *startblock
}
