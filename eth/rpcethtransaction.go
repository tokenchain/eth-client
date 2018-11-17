package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type RpcEthTransaction struct {
	Txdata
	txExtraInfo
}

type Txdata struct {
	AccountNonce string         `json:"nonce"    gencodec:"required"`
	Price        string         `json:"gasPrice"   gencodec:"required"`
	GasLimit     string         `json:"gas"      gencodec:"required"`
	Recipient    common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       string         `json:"value"    gencodec:"required"`
	Payload      []byte         `json:"input"    gencodec:"required"`

	// Signature values
	V string `json:"v" gencodec:"required"`
	R string `json:"r" gencodec:"required"`
	S string `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash common.Hash `json:"hash" rlp:"-"`
}

type txExtraInfo struct {
	BlockNumber string         `json:"blockNumber,omitempty"`
	BlockHash   common.Hash    `json:"blockHash,omitempty"`
	From        common.Address `json:"from,omitempty"`
}

func (r *RpcEthTransaction) FromAddress() common.Address {
	return r.From
}

func (r *RpcEthTransaction) ToAddress() common.Address {
	return r.Recipient
}

func (r *RpcEthTransaction) GetETHAmount() *big.Int {
	a, err := hexutil.DecodeBig(r.Amount)
	if err != nil {
		log.Info("error to decode int - ETH value")
	}
	return a
}

func (r *RpcEthTransaction) GetETHPrice() *big.Int {
	a, err := hexutil.DecodeBig(r.Price)
	if err != nil {
		log.Info("error to decode int - ETH Price")
	}
	return a
}

func (r *RpcEthTransaction) GetETHGasLimit() *big.Int {
	a, err := hexutil.DecodeBig(r.GasLimit)
	if err != nil {
		log.Info("error to decode int - ETH Gas Limit")
	}
	return a
}
