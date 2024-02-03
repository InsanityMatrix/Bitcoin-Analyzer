package main

type Wallet struct {
	Pubkey       PublicKey     `json:"address"`
	Transactions []Transaction `json:"transactions"`
	IsExchange   bool          `json:"isExchange"`
}

type PublicKey struct {
	Addresses []string `json:"addresses"`
	Type      string   `json:"type"`
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
}

func (w *Wallet) addTransaction(transaction Transaction) {
	w.Transactions = append(w.Transactions, transaction)
}
func (w *Wallet) addAddress(address string) {
	w.Pubkey.Addresses = append(w.Pubkey.Addresses, address)
}
