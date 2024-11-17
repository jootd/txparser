package transaction

import (
	"github.com/jootd/txparser/parser"
)

type Transaction struct {
	Hash        string `json:"hash"`
	BlockNumber string `json:"block_number"`
	From        string `json:"from"`
	To          string `json:"to"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gas_price"`
	Amount      string `json:"amount"`
	Nonce       string `json:"nonce"`
	Index       string `json:"index"`
}

func fromParserTransaction(tx parser.Transaction) Transaction {
	return Transaction{
		Hash:        tx.Hash,
		BlockNumber: tx.BlockNumber.String(),
		From:        tx.From,
		To:          tx.To,
		Gas:         tx.Gas.String(),
		GasPrice:    tx.GasPrice.String(),
		Amount:      tx.Amount.String(),
		Nonce:       tx.Nonce.String(),
		Index:       tx.Index.String(),
	}
}

func fromTransactionSlice(txs []parser.Transaction) []Transaction {
	handlerTxs := []Transaction{}
	for _, tx := range txs {
		handlerTxs = append(handlerTxs, fromParserTransaction(tx))
	}

	return handlerTxs
}
