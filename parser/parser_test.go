package parser

import (
	"context"
)

type MockRpcClient struct {
}

func (m *MockRpcClient) GetLatestBlockNumber(ctx context.Context) (string, error) {
	return "0x", nil
}
func (m *MockRpcClient) GetTransactions(ctx context.Context, blockNumberHex string) ([]Transaction, error) {

	return []Transaction{}, nil
}
