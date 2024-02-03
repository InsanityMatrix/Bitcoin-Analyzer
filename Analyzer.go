package main

import (
	"log"

	"github.com/toorop/go-bitcoind"
)

type BitcoinAnalyzer struct {
	CurrentBlock     uint64            `json:"currentBlock"`
	MonitorThreshold float64           `json:"monitorThreshold"`
	MonitoredWallets []MonitoredWallet `json:"monitoredWallets"`
}

func (a *BitcoinAnalyzer) updateBlock(block uint64) {
	a.CurrentBlock = block
}
func (a *BitcoinAnalyzer) getCurrentBlock() uint64 {
	return a.CurrentBlock
}

func (a *BitcoinAnalyzer) addMonitoredWallet(wallet MonitoredWallet) {
	a.MonitoredWallets = append(a.MonitoredWallets, wallet)
}
func (a *BitcoinAnalyzer) isWalletMonitored(scriptPubKey bitcoind.ScriptPubKey) (int, bool) {
	for walletIndex, wallet := range a.MonitoredWallets {
		for _, address := range wallet.Monitored.Pubkey.Addresses {
			if address == scriptPubKey.Address {
				return walletIndex, true
			}
		}
		//If address not logged, but is same pub key, add address to wallet
		if wallet.Monitored.Pubkey.Asm == scriptPubKey.Asm || wallet.Monitored.Pubkey.Hex == scriptPubKey.Hex {
			a.MonitoredWallets[walletIndex].Monitored.addAddress(scriptPubKey.Address)
			return walletIndex, true
		}
	}
	return 0, false

}

//TODO:
//Check if Transactions/Wallets are dead
//Add Monitored Transactions through wallets
//Check if Monitored Wallet is exchange
//How to handle following UTXO's?
//Saving to files
//Will this program eat all of the ram on the computer
func (a *BitcoinAnalyzer) analyzeBlock(blockNum uint64) uint64 {
	//Get the block hash of the current block
	blockhash, err := CLIENT.GetBlockHash(blockNum)
	if err != nil {
		log.Fatal(err)
	}
	//Get the block information
	block, err := CLIENT.GetBlock(blockhash)
	if err != nil {
		log.Fatal(err)
	}
	//Iterate through the transactions in the block
	for _, txid := range block.Tx {
		//Get the transaction information
		transaction, err := CLIENT.GetRawTransaction(txid, true)
		if err != nil {
			log.Fatal(err)
		}

		// Type assert transaction to the appropriate type
		tx, ok := transaction.(*bitcoind.RawTransaction)
		if !ok {
			log.Fatal("Invalid transaction type")
		}
		// Iterate through the vouts in the transaction
		for _, vout := range tx.Vout {
			// Check if the vout is a monitored address
			if vout.Value > a.MonitorThreshold {
				// Check if the address is already being monitored, if it is add TX to already monitored address
				walletIndex, isMonitored := a.isWalletMonitored(vout.ScriptPubKey)
				if isMonitored {
					// Add the transaction to the monitored address
					a.MonitoredWallets[walletIndex].Monitored.addTransaction(Transaction{
						TransactionID:    txid,
						Sender:           tx.Vin[0].ScriptSig.Hex,
						Amount:           vout.Value,
						TransactionBlock: blockNum,
					})
				} else {
					// If it is not, create a new monitored address and add the transaction to it
					monitored := MonitoredWallet{
						Monitored: Wallet{
							Pubkey: PublicKey{
								Addresses: []string{vout.ScriptPubKey.Address},
								Type:      vout.ScriptPubKey.Type,
								Asm:       vout.ScriptPubKey.Asm,
								Hex:       vout.ScriptPubKey.Hex,
							},
							Transactions: []Transaction{
								{
									TransactionID:    txid,
									Sender:           tx.Vin[0].ScriptSig.Hex,
									Amount:           vout.Value,
									TransactionBlock: blockNum,
								},
							},
						},
						SinceBlock: blockNum,
					}
					a.addMonitoredWallet(monitored)
				}

			}
		}
	}

	a.updateBlock(blockNum)
	//Now Maintain Each List

	//Check if any monitored address is dead (No transactions in 2016 blocks)
	for index, wallet := range a.MonitoredWallets {
		if wallet.isDead(blockNum) {
			//Remove the monitored address
			a.MonitoredWallets = append(a.MonitoredWallets[:index], a.MonitoredWallets[index+1:]...)
		}
	}

	return blockNum + 1
}
