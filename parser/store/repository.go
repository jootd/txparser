package store

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jootd/txparser/parser"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type DB interface {
	Get(key string) (string, bool)
	Set(key string, value string)
}

type Repository struct {
	db DB
}

func NewRepository(db DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) AddAddress(address parser.Address) error {
	key := "addrs"
	return updateDbSlice(r.db, key, string(address))
}

func (r *Repository) IsSubscribed(address parser.Address) bool {
	dbString, ok := r.db.Get("addrs")
	if !ok {
		return false
	}
	var subscribedAddrs []string
	var buf bytes.Buffer
	buf.WriteString(dbString)
	if err := gob.NewDecoder(&buf).Decode(&subscribedAddrs); err != nil {
		log.Printf("error in decoding  %s\n", err.Error())
		return false
	}
	for _, addr := range subscribedAddrs {
		if addr == string(address) {
			return true
		}
	}

	return false
}

func (r *Repository) UpdateCurrentBlock(blockHex string) {
	r.db.Set("cblock", blockHex)
}

func (r *Repository) GetCurrrentBlock() (string, error) {
	dbString, ok := r.db.Get("cblock")
	if !ok {
		return "", ErrRecordNotFound
	}
	return dbString, nil
}

func (r *Repository) AddTransactions(txs []parser.Transaction) {
	dbTxs := fromTransactionSlice(txs)
	for _, tx := range dbTxs {
		OutboundKey := fmt.Sprintf("addrs:%s:txs", tx.From)
		if err := updateDbSlice(r.db, OutboundKey, tx); err != nil {
			log.Printf("update %s   with new txs failed %s\n", OutboundKey, err.Error())
		}

		InboundKey := fmt.Sprintf("addrs:%s:txs", tx.To)
		if err := updateDbSlice(r.db, InboundKey, tx); err != nil {
			log.Printf("update %s  with new txs failed  %s\n", OutboundKey, err.Error())
		}
	}
}

func (r *Repository) GetTransactionsBy(addr parser.Address) []parser.Transaction {
	key := fmt.Sprintf("addrs:%s:txs", strings.ToLower(string(addr)))
	dbString, exists := r.db.Get(key)
	if !exists {
		log.Printf("key not exists %s", key)
		return []parser.Transaction{}
	}

	var dbTxs []dbTransaction
	var buf bytes.Buffer

	buf.WriteString(dbString)

	if err := gob.NewDecoder(&buf).Decode(&dbTxs); err != nil {
		log.Printf("error in decoding  %s\n", err.Error())
		return []parser.Transaction{}
	}

	log.Printf("getting dbTsxs length %d", len(dbTxs))

	return toTransactionSlice(dbTxs)
}

func updateDbSlice[T any](db DB, key string, value T) error {
	dbSliceString, exists := db.Get(key)
	if !exists {
		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode([]T{value}); err != nil {
			return err
		}

		db.Set(key, buf.String())
		return nil
	}

	var existingSlice []T
	var buf bytes.Buffer

	buf.WriteString(dbSliceString)

	if err := gob.NewDecoder(&buf).Decode(&existingSlice); err != nil {
		return fmt.Errorf("failed to decode existing data: %w", err)
	}

	existingSlice = append(existingSlice, value)

	var updatedBuf bytes.Buffer
	if err := gob.NewEncoder(&updatedBuf).Encode(existingSlice); err != nil {
		return err
	}

	db.Set(key, updatedBuf.String())
	return nil
}
