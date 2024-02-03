package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DataDir struct {
	dir string
}

type Address struct {
	Owner          string        `json:"owner"`
	Address        string        `json:"address"`
	Transactions   []Transaction `json:"transactions"`
	CurrentBalance float64       `json:"currentBalance"`
}
type Transaction struct {
	Sender           string  `json:"sender"`
	Recipient        string  `json:"recipient"`
	Amount           float64 `json:"amount"`
	TransactionID    string  `json:"transactionID"`
	TransactionBlock uint64  `json:"transactionBlock"`
}

var DATADIR DataDir

//Simple Check if DataDir exists
func (d DataDir) exists() bool {
	_, err := os.Stat(d.dir)
	return !os.IsNotExist(err)
}

//Create DataDir
func (d DataDir) create() error {
	return os.MkdirAll(d.dir, 0755)
}

//Create a file in DataDir
func (d DataDir) createFile(fname string) error {
	_, err := os.Create(d.dir + "/" + fname)
	return err
}

//Create a dir in DataDir
func (d DataDir) createDir(fname string) error {
	return os.MkdirAll(d.dir+string(filepath.Separator)+fname, 0755)
}

func (d DataDir) readConfig() (BitcoinAnalyzer, error) {
	config, _ := ioutil.ReadFile(d.dir + string(filepath.Separator) + "config.json")
	var analyzer BitcoinAnalyzer
	err := json.Unmarshal(config, &analyzer)
	return analyzer, err
}

func (d DataDir) writeConfig() error {
	config, _ := json.Marshal(ANALYZER)
	return ioutil.WriteFile(d.dir+string(filepath.Separator)+"config.json", config, 0644)
}
