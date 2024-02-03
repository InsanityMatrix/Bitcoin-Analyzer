package main

import (
	"log"
	"os"
	"path/filepath"

	bitcoind "github.com/toorop/go-bitcoind"
)

const (
	SERVER_HOST       = "127.0.0.1"
	SERVER_PORT       = 8332
	USER              = "bitcoin_analyzer"
	USESSL            = false
	WALLET_PASSPHRASE = "WalletPassphrase"
)

var ANALYZER BitcoinAnalyzer
var CLIENT *bitcoind.Bitcoind

func main() {
	DATADIRstr, STARTBLOCK := getArgs()
	DATADIRpath, _ := filepath.Abs(DATADIRstr)
	DATADIR = DataDir{DATADIRpath}
	//Connect to btc rpc server
	password := os.Getenv("ANALYZER_PASS")
	CLIENT, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, password, USESSL)
	if err != nil {
		log.Fatal(err)
	}

	existed := DATADIR.exists()
	if !existed {
		initialize(STARTBLOCK)
	} else {
		//Get the latest block from config
		ANALYZER, err = DATADIR.readConfig()
	}

	//Get the current block count()
	blockcount, err := CLIENT.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Current Block count: %d", blockcount)
	//IF current blockcount is greater than the blockcount in the config file, start analyzing blocks

}

func initialize(STARTBLOCK uint64) {
	//Check if Datadir exists - If not create it
	if !DATADIR.exists() {
		//Create Datadir
		err := DATADIR.create()
		if err != nil {
			log.Fatal(err)
		}
	}

	//Create the necessary folders
	err := DATADIR.createDir("addresses")
	if err != nil {
		log.Fatal(err)
	}
	err = DATADIR.createDir("blocks")
	if err != nil {
		log.Fatal(err)
	}
	err = DATADIR.createDir("transactions")
	if err != nil {
		log.Fatal(err)
	}

	//Create a config file to keep track of blocks, transactions and addresses
	err = DATADIR.createFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	if STARTBLOCK == 0 {
		blockcount, err := CLIENT.GetBlockCount()
		if err != nil {
			log.Fatal(err)
		}
		ANALYZER = BitcoinAnalyzer{CurrentBlock: blockcount - 1}

	}
	//Write to config file the current BitcoinAnalyzer struct
	err = DATADIR.writeConfig()
	if err != nil {
		log.Fatal(err)
	}
}
