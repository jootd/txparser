package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jootd/txparser/parser"
)

type Config struct {
	BaseURl  string
	Version  string
	ClientId int
}

type Client struct {
	httpClient *http.Client
	baseUrl    string
	clientId   int
	version    string
}

func NewClient(opts Config, client *http.Client) *Client {
	return &Client{
		httpClient: client,
		baseUrl:    opts.BaseURl,
		clientId:   opts.ClientId,
		version:    opts.Version,
	}
}

func (c *Client) methodBlockNumber() string   { return "eth_blockNumber" }
func (c *Client) methodBlockByNumber() string { return "eth_getBlockByNumber" }

func (c *Client) GetLatestBlockNumber(ctx context.Context) (string, error) {
	reqBody := RPCRequest{
		JSONRPC: c.version,
		Method:  c.methodBlockNumber(),
		Params:  []interface{}{},
		ID:      c.clientId,
	}

	res, err := c.makeRequest(ctx, "POST", reqBody)
	if err != nil {
		return "", fmt.Errorf("client get latest block number: %w", err)
	}
	return res.Result.(string), nil
}

func (c *Client) GetTransactions(ctx context.Context, blockNumberHex string) ([]parser.Transaction, error) {
	reqBod := RPCRequest{
		JSONRPC: c.version,
		Method:  c.methodBlockByNumber(),
		Params:  []interface{}{blockNumberHex, true},
	}

	res, err := c.makeRequest(ctx, "POST", reqBod)
	if err != nil {
		return []parser.Transaction{}, fmt.Errorf("client: %w", err)
	}

	var blockResponse BlockResponse
	resultBytes, err := json.Marshal(res.Result)
	if err != nil {
		return []parser.Transaction{}, fmt.Errorf("client: marshall, %w", err)
	}

	if err := json.Unmarshal(resultBytes, &blockResponse); err != nil {
		return []parser.Transaction{}, fmt.Errorf("client: unmarshall, %w", err)
	}
	return toTransactionSlice(blockResponse), nil
}

func (c *Client) makeRequest(ctx context.Context, method string, reqBody RPCRequest) (RPCResponse, error) {
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return RPCResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseUrl, bytes.NewBuffer(reqBytes))
	if err != nil {
		return RPCResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RPCResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RPCResponse{}, err
	}
	var res RPCResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return RPCResponse{}, err
	}

	return res, nil
}
