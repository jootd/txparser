package main

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jootd/txparser/cmd/handlers/transaction"
	"github.com/jootd/txparser/parser"
	"github.com/jootd/txparser/parser/client/rpc"
	"github.com/jootd/txparser/parser/store"
	"github.com/jootd/txparser/parser/store/db"
)

func main() {
	log.SetOutput(os.Stdout)

	db := db.NewMemoryStorage()

	store := store.NewRepository(db)
	rpcClient := rpc.NewClient(rpc.Config{
		BaseURl:  "https://ethereum-rpc.publicnode.com",
		Version:  "2.0",
		ClientId: 83,
	}, http.DefaultClient)

	parser, err := parser.NewTxParser(
		store,
		rpcClient,
		big.NewInt(21207502),
	)
	if err != nil {
		log.Fatal(err)
	}

	go parser.Start()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	handler := transaction.NewHandler(parser)
	http.HandleFunc("/current-block", handler.GetCurrentBlockHandler)
	http.HandleFunc("/subscribe", handler.SubscribeHandler)
	http.HandleFunc("/transactions", handler.GetTransactionsHandler)

	server := http.Server{
		Addr: ":8080",
	}

	go func() {
		log.Println("starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server: %v", err)
		}
	}()

	<-shutdown
	parser.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutting down server: %v", err)
	}

	log.Println("server shut down successfully")
}
