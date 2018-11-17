package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"encoding/json"
)

type TokenBalance struct {
	Contract       common.Address
	Wallet         common.Address
	Name           string
	Symbol         string
	Balance        *big.Int
	ETH            *big.Int
	Decimals       int64
	Block          int64
	InternalUserId int64
}

func (tb *TokenBalance) ETHString() string {
	return bigIntString(tb.ETH, 18)
}

func (tb *TokenBalance) BalanceString() string {
	if tb.Decimals == 0 {
		return tb.Balance.String()
	}
	return bigIntString(tb.Balance, tb.Decimals)
}

func (tb *TokenBalance) ToJSON() string {
	jsonData := TokenBalanceJson{
		Contract: tb.Contract.String(),
		Wallet:   tb.Wallet.String(),
		Name:     tb.Name,
		Symbol:   tb.Symbol,
		Balance:  tb.BalanceString(),
		ETH:      tb.ETHString(),
		Decimals: tb.Decimals,
		Block:    tb.Block,
	}
	d, _ := json.Marshal(jsonData)
	return string(d)
}
