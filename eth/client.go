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
	"errors"

	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/p2p"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum"
)

// client defines typed wrappers for the Ethereum RPC API.
type ClientTokenEth struct {
	*ethclient.Client
	rpc *ethrpc.Client
}

type RpcEthTransaction struct {
	Data Txdata
	txExtraInfo
}

type Txdata struct {
	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64          `json:"gas"      gencodec:"required"`
	Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    gencodec:"required"`
	Payload      []byte          `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*ClientTokenEth, error) {
	rpc, err := ethrpc.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(rpc), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(rpc *ethrpc.Client) *ClientTokenEth {
	return &ClientTokenEth{
		Client: ethclient.NewClient(rpc),
		rpc:    rpc,
	}
}

// Close closes an existing RPC connection.
func (c *ClientTokenEth) Close() {
	c.rpc.Close()
}

// Retrieve list of APIs available on the server
func (c *ClientTokenEth) SupportedModules() (map[string]string, error) {
	r, err := c.rpc.SupportedModules()
	return r, err
}

// ----------------------------------------------------------------------------
// eth

// SendRawTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (c *ClientTokenEth) SendRawTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.SendTransaction(ctx, tx)
}

// BlockNumber returns the current block number.
func (c *ClientTokenEth) BlockNumber(ctx context.Context) (*big.Int, error) {
	var r string
	err := c.rpc.CallContext(ctx, &r, "eth_blockNumber")
	if err != nil {
		return nil, err
	}
	h, err := hexutil.DecodeBig(r)
	return h, err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

// ----------------------------------------------------------------------------
// admin

// AddPeer connects to the given nodeURL.
func (c *ClientTokenEth) AddPeer(ctx context.Context, nodeURL string) error {
	var r bool
	// TODO: Result needs to be verified
	err := c.rpc.CallContext(ctx, &r, "admin_addPeer", nodeURL)
	if err != nil {
		return err
	}
	return err
}

// AdminPeers returns the number of connected peers.
func (c *ClientTokenEth) AdminPeers(ctx context.Context) ([]*p2p.PeerInfo, error) {
	var r []*p2p.PeerInfo
	// The response data type are bytes, but we cannot parse...
	err := c.rpc.CallContext(ctx, &r, "admin_peers")
	if err != nil {
		return nil, err
	}
	return r, err
}

// NodeInfo gathers and returns a collection of metadata known about the host.
func (c *ClientTokenEth) NodeInfo(ctx context.Context) (*p2p.PeerInfo, error) {
	var r *p2p.PeerInfo
	err := c.rpc.CallContext(ctx, &r, "admin_nodeInfo")
	if err != nil {
		return nil, err
	}
	return r, err
}

// ----------------------------------------------------------------------------
// miner

// SetMiningAccount sets etherbase
func (c *ClientTokenEth) SetMiningAccount(ctx context.Context, account string) error {
	etherbase := common.HexToAddress(account)
	var r bool
	// TODO: Result needs to be verified
	err := c.rpc.CallContext(ctx, &r, "miner_setEtherbase", etherbase)
	if err != nil {
		return err
	}
	return err
}

// StartMining starts mining operation.
func (c *ClientTokenEth) StartMining(ctx context.Context) error {
	var r []byte
	// TODO: Result needs to be verified
	// The response data type are bytes, but we cannot parse...
	err := c.rpc.CallContext(ctx, &r, "miner_start", nil)
	if err != nil {
		return err
	}
	return err
}

// StopMining stops mining.
func (c *ClientTokenEth) StopMining(ctx context.Context) error {
	err := c.rpc.CallContext(ctx, nil, "miner_stop", nil)
	if err != nil {
		return err
	}
	return err
}

// Generic client.Client functions
func (c *ClientTokenEth) GetInfo(ctx context.Context) (string, error) {
	type info struct {
		NodeInfo    *p2p.PeerInfo   `json:"nodeInfo"`
		AdminPeers  []*p2p.PeerInfo `json:"adminPeers"`
		BlockNumber string          `json:"blockNumber"`
	}
	block, err := c.BlockNumber(ctx)
	if err != nil {
		return "", err
	}
	resp := &info{BlockNumber: block.String()}

	ni, err := c.NodeInfo(ctx)
	if err == nil {
		resp.NodeInfo = ni
	}

	ap, err := c.AdminPeers(ctx)
	if err == nil {
		resp.AdminPeers = ap
	}

	out, err := json.Marshal(resp)
	return string(out), err
}

func (c *ClientTokenEth) GenerateKey(ctx context.Context) (address, private string, err error) {
	key, _ := crypto.GenerateKey()
	address = crypto.PubkeyToAddress(key.PublicKey).Hex()
	private = hex.EncodeToString(key.D.Bytes())
	return
}

// todo: SendERC-20 also
// https://ethereum.stackexchange.com/questions/10486/raw-transaction-data-in-go
func privateKeyToPubAddress(privKey *ecdsa.PrivateKey) (*common.Address, error) {
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}
	pubAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
	return &pubAddress, nil
}

func (c *ClientTokenEth) SendAmount(ctx context.Context, fromPriv, toPub, amount string) error {
	fromPrivKey, err := crypto.HexToECDSA(fromPriv)
	if err != nil {
		return err
	}

	// convert fromPrivKey to fromPubAddress
	fromPubAddress, err := privateKeyToPubAddress(fromPrivKey)
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress(toPub)

	amountInt := new(big.Int)
	amountInt.SetString(amount, 10)

	// find a gas price
	gasPrice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	// find pending nonce for from account
	nonce, err := c.PendingNonceAt(ctx, *fromPubAddress)
	if err != nil {
		return err
	}

	// assume gas limit of 21,000
	tx := types.NewTransaction(nonce, toAddress, amountInt,
		21000, gasPrice, nil)

	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, fromPrivKey)
	if err != nil {
		return err
	}

	return c.SendTransaction(ctx, signTx)
}

func (c *ClientTokenEth) GetBalance(ctx context.Context, account string) (string, error) {
	address := common.HexToAddress(account)
	// at nil means last known balance
	balance, err := c.BalanceAt(ctx, address, nil)
	if err != nil {
		return "", err
	}
	return balance.Text(10), nil
}

func (c *ClientTokenEth) LatestConfirmedTransactionCount(ctx context.Context) (uint, error) {
	var num hexutil.Uint
	err := c.rpc.CallContext(ctx, &num, "eth_getBlockTransactionCountByNumber", "latest")
	return uint(num), err
}

func (c *ClientTokenEth) TransactionCountByBlockNumber(ctx context.Context, number *big.Int) (uint, error) {
	var num hexutil.Uint
	err := c.rpc.CallContext(ctx, &num, "eth_getBlockTransactionCountByNumber", toBlockNumArg(number))
	return uint(num), err
}

func (c *ClientTokenEth) TransactionByBlockNumberIndex(ctx context.Context, number *big.Int, index *big.Int) (*RpcEthTransaction, error) {
	var json *RpcEthTransaction
	err := c.rpc.CallContext(ctx, &json, "eth_getTransactionByBlockNumberAndIndex", toBlockNumArg(number))
	if err == nil {
		if json == nil {
			return nil, ethereum.NotFound
		}
	}
	//if json.From != nil && json.BlockHash != nil {
	//setSenderFromServer(json.tx, *json.From, *json.BlockHash)
	//}
	return json, err
}
