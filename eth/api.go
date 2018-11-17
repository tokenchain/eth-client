// Copyright 2017 AMIS Technologies
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/p2p"
)

type Client interface {
	//	generic.Client

	// eth
	BlockNumber(ctx context.Context) (*big.Int, error)
	SendRawTransaction(ctx context.Context, tx *types.Transaction) error

	// admin
	AddPeer(ctx context.Context, nodeURL string) error
	AdminPeers(ctx context.Context) ([]*p2p.PeerInfo, error)
	NodeInfo(ctx context.Context) (*p2p.PeerInfo, error)
	SupportedModules() (map[string]string, error)

	// miner
	StartMining(ctx context.Context) error
	StopMining(ctx context.Context) error
	SetMiningAccount(ctx context.Context, account string) error

	// eth client
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error)
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	LatestConfirmedTransactionCount(ctx context.Context) (uint, error)
	TransactionCountByBlockNumber(ctx context.Context, number *big.Int) (uint, error)
	TransactionByBlockNumberIndex(ctx context.Context, number *big.Int, index *big.Int) (*RpcEthTransaction, error)
	SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	NetworkID(ctx context.Context) (*big.Int, error)
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error)
	PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	PendingTransactionCount(ctx context.Context) (uint, error)
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	GetTokenBalance(ctx context.Context, token_contract common.Address, account_wallet common.Address, blockNumber *big.Int) (*TokenBalance, error)
	GetTokenBalanceLatest(ctx context.Context, token_contract common.Address, account_wallet common.Address) (*TokenBalance, error)
}


// Package tokenbalance is used to fetch the latest token balance for any Ethereum address and ERC20 token. You can install
// this package/CLI or you can use basic HTTP GET on the public TokenBalance server.
//
// Mainnet API Endpoint:
// https://api.tokenbalance.com
//
// Example: https://api.tokenbalance.com/balance/0xa74476443119A942dE498590Fe1f2454d7D4aC0d/0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a
// Response: `5401731.086778292432427406`
//
// Example: https://api.tokenbalance.com/token/0xa74476443119A942dE498590Fe1f2454d7D4aC0d/0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a
// Response:
// ```
// {
// "token": "0xa74476443119A942dE498590Fe1f2454d7D4aC0d",
// "wallet": "0xda0AEd568D9A2dbDcBAFC1576fedc633d28EEE9a",
// "name": "Golem Network token",
// "symbol": "GNT",
// "balance": "5401731.086778292432427406",
// "eth_balance": "0.985735366999999973",
// "decimals": 18,
// "block": 6461672
// }
// ```
//
// Ropsten Testnet API Endpoint:
// https://test.tokenbalance.com
//
// Rinkeby Testnet API Endpoint:
// https://rinkeby.tokenbalance.com
//
// More info on: https://github.com/hunterlong/tokenbalance
