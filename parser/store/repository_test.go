package store

import (
	"bytes"
	"encoding/gob"
	"slices"
	"testing"

	"github.com/jootd/txparser/parser"
	"github.com/jootd/txparser/parser/store/db"
)

func TestRepository_AddAddress(t *testing.T) {
	memDB := db.NewMemoryStorage()
	repo := NewRepository(memDB)

	address := parser.Address("0x123")
	err := repo.AddAddress(address)

	dbString, exists := memDB.Get("addrs")
	if !exists {
		t.Fatalf("")
	}

	var addrs []string
	buf := bytes.NewBufferString(dbString)
	err = gob.NewDecoder(buf).Decode(&addrs)
	if err != nil {
		t.Fatal("")

	}
	if !slices.Contains(addrs, string(address)) {
		t.Fatal("does not contains")

	}

}

func TestRepository_IsSubscribed(t *testing.T) {
	memDB := db.NewMemoryStorage()
	repo := NewRepository(memDB)

	address := parser.Address("0x123")
	_ = repo.AddAddress(address)
	if !repo.IsSubscribed(address) {
		t.Fatal("")
	}

	if repo.IsSubscribed(parser.Address("0xxx")) {
		t.Fatal("should not be there ")
	}
}

func TestRepository_UpdateGetCurrentBlock(t *testing.T) {
	memDB := db.NewMemoryStorage()
	repo := NewRepository(memDB)
	current := "0x123"
	repo.UpdateCurrentBlock(current)

	value, err := repo.GetCurrrentBlock()
	if err != nil {
		t.Error(err)
		return
	}

	if value != current {
		t.Fatal("does not equal")
	}
}

func TestRepository_AddGetTransactions(t *testing.T) {
	memDB := db.NewMemoryStorage()
	repo := NewRepository(memDB)

	fromAddr := "0x123"
	toAddr1 := "0x456"
	toAddr2 := "0x789"
	txs := []parser.Transaction{
		{From: fromAddr, To: toAddr1, Hash: "0x"},
		{From: fromAddr, To: toAddr2, Hash: "0xx"},
	}
	repo.AddTransactions(txs)

	txsFromDb := repo.GetTransactionsBy(parser.Address(fromAddr))
	if len(txsFromDb) != 2 {
		t.Fatalf("txs  length for  %s, expected %d , got %d", toAddr2, 2, len(txsFromDb))
	}

	txsTo1DB := repo.GetTransactionsBy(parser.Address(toAddr1))
	if len(txsTo1DB) != 1 {
		t.Fatalf("txs  length for  %s, want %d , got %d", toAddr1, 1, len(txsTo1DB))
	}

	txsTo2DB := repo.GetTransactionsBy(parser.Address(toAddr1))
	if len(txsTo2DB) != 1 {
		t.Fatalf("txs  length for  %s, want %d , got %d", toAddr2, 1, len(txsTo2DB))
	}
}
