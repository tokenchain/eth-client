package eth

import (
	"math/big"
	"fmt"
	"bytes"
	"github.com/ethereum/go-ethereum/common"
)

func GetSymbolFromId(contract_address common.Address) string {
	for symbol, address := range token_list {
		//	fmt.Println("Key:", key, "=>", "Element:", element)
		if contract_address.String() == address {
			return symbol
		}
	}
	return "MISSING"
}
func GetContractFromSymbol(sym string) common.Address {
	for symbol, address := range token_list {
		//	fmt.Println("Key:", key, "=>", "Element:", element)
		if symbol == sym {
			return common.HexToAddress(address)
		}
	}
	return common.HexToAddress("")
}
func symbolFix(contract string) string {
	for symbol, address := range token_list {
		//	fmt.Println("Key:", key, "=>", "Element:", element)
		if common.HexToAddress(contract).String() == address {
			return symbol
		}
	}
	return "MISSING"
}

func bigIntString(balance *big.Int, decimals int64) string {
	amount := bigIntFloat(balance, decimals)
	deci := fmt.Sprintf("%%0.%vf", decimals)
	return clean(fmt.Sprintf(deci, amount))
}

func bigIntFloat(balance *big.Int, decimals int64) *big.Float {
	if balance.Sign() == 0 {
		return big.NewFloat(0)
	}
	bal := big.NewFloat(0)
	bal.SetInt(balance)
	pow := bigPow(10, decimals)
	p := big.NewFloat(0)
	p.SetInt(pow)
	bal.Quo(bal, p)
	return bal
}

func bigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}

func clean(newNum string) string {
	stringBytes := bytes.TrimRight([]byte(newNum), "0")
	newNum = string(stringBytes)
	if stringBytes[len(stringBytes)-1] == 46 {
		newNum += "0"
	}
	if stringBytes[0] == 46 {
		newNum = "0" + newNum
	}
	return newNum
}
