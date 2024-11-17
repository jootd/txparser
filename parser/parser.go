package parser

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/jootd/txparser/sdk"
)

type Storer interface {
	AddAddress(address Address) error
	UpdateCurrentBlock(blockHex string)
	GetCurrrentBlock() (string, error)
	GetTransactionsBy(address Address) []Transaction
	AddTransactions(txs []Transaction)
	IsSubscribed(address Address) bool
}

type Client interface {
	GetLatestBlockNumber(ctx context.Context) (string, error)
	GetTransactions(ctx context.Context, blockNumberHex string) ([]Transaction, error)
}

type TxParser struct {
	store     Storer
	client    Client
	fromBlock *big.Int
	quit      chan struct{}
}

func NewTxParser(store Storer, client Client, fromBlock *big.Int) (*TxParser, error) {
	if fromBlock == nil {
		return &TxParser{}, errors.New("start block was not provided")
	}
	return &TxParser{
		quit:      make(chan struct{}),
		store:     store,
		client:    client,
		fromBlock: fromBlock,
	}, nil
}

func (p *TxParser) Start() {
	for {
		select {
		case <-p.quit:
			log.Println("shutting down parser...")
			return
		default:

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			latestBlockhex, err := p.client.GetLatestBlockNumber(ctx)
			cancel()
			if err != nil {
				log.Printf("error fetching latest block: %v\n", err)
				return
			}

			latestBlockInt := sdk.FromHex(latestBlockhex)
			for p.fromBlock.Cmp(latestBlockInt) <= 0 {

				select {
				case <-p.quit:
					log.Println("Shutting down parser")
					return
				default:
					txCtx, txCancel := context.WithTimeout(context.Background(), time.Second*10)
					transactions, err := p.client.GetTransactions(txCtx, sdk.ToHex(p.fromBlock))
					txCancel()
					if err != nil {
						log.Printf("failed to get latest block number, in block %d, %s\n", p.fromBlock, err.Error())
						continue
					}
					p.store.AddTransactions(transactions)
					// increment
					p.store.UpdateCurrentBlock(sdk.ToHex(p.fromBlock))
					p.fromBlock.Add(p.fromBlock, big.NewInt(1))

				}
			}
		}
	}
}

func (p *TxParser) Close() {
	p.quit <- struct{}{}
}

// cblk stands for current block
func (p *TxParser) GetCurrentBlock() *big.Int {
	currBlock, err := p.store.GetCurrrentBlock()
	if err != nil {
		log.Printf("parser get current block:  %s", err.Error())
		return big.NewInt(0)
	}

	return sdk.FromHex(currBlock)
}

func (txp *TxParser) Subscribe(address string) bool {
	addr, err := parseAddress(address)
	if err != nil {
		log.Printf("parser subscribe: %s", err.Error())
		return false
	}
	txp.store.AddAddress(addr)
	return true
}

func (p *TxParser) GetTransactions(address string) []Transaction {
	addr, err := parseAddress(address)
	if err != nil {
		log.Printf("parser validation :%s", err.Error())
		return []Transaction{}
	}

	if !p.store.IsSubscribed(addr) {
		return []Transaction{}
	}

	return p.store.GetTransactionsBy(addr)

}
