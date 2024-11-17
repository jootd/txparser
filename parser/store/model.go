package store

import (
	"github.com/jootd/txparser/parser"
	"github.com/jootd/txparser/sdk"
)

type dbTransaction struct {
	Hash        string
	BlockNumber string
	From        string
	To          string
	Gas         string
	GasPrice    string
	Amount      string
	Nonce       string
	Index       string
}

func fromTransaction(tx parser.Transaction) dbTransaction {
	return dbTransaction{
		Hash:        tx.Hash,
		BlockNumber: sdk.ToHex(tx.BlockNumber),
		From:        tx.From,
		To:          tx.To,
		Gas:         sdk.ToHex(tx.Gas),
		GasPrice:    sdk.ToHex(tx.GasPrice),
		Amount:      sdk.ToHex(tx.Amount),
		Nonce:       sdk.ToHex(tx.Nonce),
		Index:       sdk.ToHex(tx.Index),
	}
}

func fromTransactionSlice(txs []parser.Transaction) []dbTransaction {
	var dbTxs []dbTransaction
	for _, tx := range txs {
		dbTxs = append(dbTxs, fromTransaction(tx))
	}
	return dbTxs
}

func toTransaction(dbTx dbTransaction) parser.Transaction {
	tx := parser.Transaction{
		Hash:        dbTx.Hash,
		BlockNumber: sdk.FromHex(dbTx.BlockNumber),
		From:        dbTx.From,
		To:          dbTx.To,
		Gas:         sdk.FromHex(dbTx.Gas),
		GasPrice:    sdk.FromHex(dbTx.GasPrice),
		Amount:      sdk.FromHex(dbTx.Amount),
		Nonce:       sdk.FromHex(dbTx.Nonce),
		Index:       sdk.FromHex(dbTx.Index),
	}

	return tx
}

func toTransactionSlice(dbTx []dbTransaction) []parser.Transaction {
	var txs []parser.Transaction
	for _, tx := range dbTx {
		txs = append(txs, toTransaction(tx))
	}

	return txs
}
