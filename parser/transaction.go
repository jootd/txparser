package parser

import (
	"math/big"
)

// type TxType string
// var (
// 	Inbound  TxType = "Inbound"
// 	Outbound TxType = "Outbound"
// )

type Transaction struct {
	Hash        string
	BlockNumber *big.Int
	From        string
	To          string
	Gas         *big.Int
	GasPrice    *big.Int
	Amount      *big.Int
	Nonce       *big.Int
	Index       *big.Int
}

// func (a Transaction) isRelatedTo(addr Address) bool {
// 	return a.From == string(addr) || a.To == string(addr)
// }

// func (a Transaction) typeFor(addr string) TxType {
// 	if a.To == addr {
// 		return Inbound
// 	}
// 	return Outbound
// }
