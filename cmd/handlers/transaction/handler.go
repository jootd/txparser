package transaction

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/jootd/txparser/parser"
)

type Parser interface {
	// last parsed block
	GetCurrentBlock() *big.Int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []parser.Transaction
}

type Handler struct {
	parser Parser
}

func NewHandler(parser Parser) *Handler {
	return &Handler{parser: parser}
}

func (h *Handler) GetCurrentBlockHandler(w http.ResponseWriter, r *http.Request) {
	block := h.parser.GetCurrentBlock()
	resp := struct {
		CurrentBlock string `json:"currentBlock"`
	}{CurrentBlock: block.String()}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *Handler) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	ok := h.parser.Subscribe(address)
	resp := struct {
		Subscribed bool `json:"subscribed"`
	}{Subscribed: ok}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *Handler) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address  is required", http.StatusBadRequest)
		return
	}

	txs := h.parser.GetTransactions(address)
	mappedTxs := fromTransactionSlice(txs)

	resp := struct {
		Transactions []Transaction `json:"transactions"`
	}{
		Transactions: mappedTxs,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
