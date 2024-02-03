package main

import (
	"log"
	"path/filepath"

	bitcoind "github.com/toorop/go-bitcoind"
)

const (
	SERVER_HOST       = "127.0.0.1"
	SERVER_PORT       = 8332
	USER              = "bitcoin_analyzer"
	PASSWD            = "!LiNK3d-Up14&iDont!TRU5T.U"
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
	CLIENT, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
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
		ANALYZER = BitcoinAnalyzer{blockcount - 1}

	}
	//Write to config file the current BitcoinAnalyzer struct
	err = DATADIR.writeConfig()
	if err != nil {
		log.Fatal(err)
	}

}
