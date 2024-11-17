# Ethereum Transaction Parser

A Go-based application that fetches the latest Ethereum block numbers and transactions, stores them in memory, and provides an HTTP interface to access transaction data and block information. This parser interacts with the Ethereum blockchain using RPC requests.

## Features

- Fetches the latest block number from the Ethereum network.
- Retrieves transactions for each block and stores them.
- Provides an HTTP API for interacting with transaction data.

## Architecture
- **TxParser**: The core component that connects to an Ethereum RPC client, fetches block numbers and transactions, and stores them in a repository.
- **RPC Client**: Communicates with the Ethereum network via RPC requests to fetch the latest block and transaction data.
- **Repository**: Stores the fetched transactions in DB (in this case in MemoryStorage).
- **MemoryStorage** Thread safe in-memory storage.
- **HTTP Server**: Exposes an HTTP interface to retrieve current block information and transaction details for given address.

## Data Storage Strategy

The application stores all retrieved transactions. Although this method might not be the most efficient in terms of memory usage and performance, it is simple and ensures that the transaction history of all addresses is preserved. This makes the system more flexible, as new addresses always have their full transaction history available.
