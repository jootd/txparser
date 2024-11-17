package rpc

import (
	"github.com/jootd/txparser/parser"
	"github.com/jootd/txparser/sdk"
)

type BlockResponse struct {
	BaseFeePerGas         string `json:"baseFeePerGas"`
	BlobGasUsed           string `json:"blobGasUsed"`
	Difficulty            string `json:"difficulty"`
	ExcessBlobGas         string `json:"excessBlobGas"`
	ExtraData             string `json:"extraData"`
	GasLimit              string `json:"gasLimit"`
	GasUsed               string `json:"gasUsed"`
	Hash                  string `json:"hash"`
	LogsBloom             string `json:"logsBloom"`
	Miner                 string `json:"miner"`
	MixHash               string `json:"mixHash"`
	Nonce                 string `json:"nonce"`
	Number                string `json:"number"`
	ParentBeaconBlockRoot string `json:"parentBeaconBlockRoot"`
	ParentHash            string `json:"parentHash"`
	ReceiptsRoot          string `json:"receiptsRoot"`
	Sha3Uncles            string `json:"sha3Uncles"`
	Size                  string `json:"size"`
	StateRoot             string `json:"stateRoot"`
	Timestamp             string `json:"timestamp"`
	TotalDifficulty       string `json:"totalDifficulty"`
	Transactions          []struct {
		AccessList []struct {
			Address     string   `json:"address"`
			StorageKeys []string `json:"storageKeys"`
		} `json:"accessList,omitempty"`
		BlockHash            string   `json:"blockHash"`
		BlockNumber          string   `json:"blockNumber"`
		ChainId              string   `json:"chainId"`
		From                 string   `json:"from"`
		Gas                  string   `json:"gas"`
		GasPrice             string   `json:"gasPrice"`
		Hash                 string   `json:"hash"`
		Input                string   `json:"input"`
		MaxFeePerGas         string   `json:"maxFeePerGas,omitempty"`
		MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas,omitempty"`
		Nonce                string   `json:"nonce"`
		R                    string   `json:"r"`
		S                    string   `json:"s"`
		To                   string   `json:"to"`
		TransactionIndex     string   `json:"transactionIndex"`
		Type                 string   `json:"type"`
		V                    string   `json:"v"`
		Value                string   `json:"value"`
		YParity              string   `json:"yParity,omitempty"`
		BlobVersionedHashes  []string `json:"blobVersionedHashes,omitempty"`
		MaxFeePerBlobGas     string   `json:"maxFeePerBlobGas,omitempty"`
	} `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
	Withdrawals      []struct {
		Address        string `json:"address"`
		Amount         string `json:"amount"`
		Index          string `json:"index"`
		ValidatorIndex string `json:"validatorIndex"`
	} `json:"withdrawals"`
	WithdrawalsRoot string `json:"withdrawalsRoot"`
}

func toTransactionSlice(block BlockResponse) []parser.Transaction {
	txs := []parser.Transaction{}
	for _, tx := range block.Transactions {
		txs = append(txs, parser.Transaction{
			Hash:        tx.Hash,
			BlockNumber: sdk.FromHex(tx.BlockNumber),
			From:        tx.From,
			To:          tx.To,
			Gas:         sdk.FromHex(tx.Gas),
			GasPrice:    sdk.FromHex(tx.GasPrice),
			Amount:      sdk.FromHex(tx.Value),
			Nonce:       sdk.FromHex(tx.Nonce),
			Index:       sdk.FromHex(tx.TransactionIndex),
		})
	}
	return txs
}
