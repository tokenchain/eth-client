package eth

import (
	"testing"
	"math/big"
	"context"
	"fmt"
	mx "github.com/tokenchain/eth-client/eth"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	geth_rpc            *mx.ClientTokenEth
	latest_block_number int64
)

func ConnectGethRpc() (client_tx *mx.ClientTokenEth) {
	location := fmt.Sprintf("http://%s:%s", "104.199.155.7", "8200")
	client_tx, err := mx.Dial(location)
	if err != nil {
		fmt.Println("Failed to dial, url: ", location, ", err: ", err)
		return
	}

	return
}

/*
	"blockHash": "0xbf87cbdf6d36b734def2c2b3770d735b56dcafda54f296955531cebdc6065d0b",
	"blockNumber": "0x25b096",
	"from": "0x3fd3adba69955f85bc34860b77a64c2c52c981ea",
	"gas": "0x30d40",
	"gasPrice": "0x4a817c800",
	"hash": "0x7e417c34c2360f8dda2dd79026fe499b299306bc2d067f0626723d00a34914bf",
	"input": "0x",
	"nonce": "0x4",
	"to": "0xc1a2f88be90d08e6899e3cd2e9021e95c3ff6324",
	"transactionIndex": "0x5",
	"value": "0x3a39735e5d31ea00",
	"v": "0x1b",
	"r": "0x70b5f34b07c090689297190814533b6f911470496bede67f6ee555b72467b228",
	"s": "0x39e2149fb6516f13ae48ca6ba256420379efa75ebd83cb52e508992af4f9ad97"

 */
func TestClientTokenEth_LatestConfirmedTransactionCount(t *testing.T) {
	geth_rpc = ConnectGethRpc()
	tx, _ := geth_rpc.TransactionByBlockNumberIndex(context.Background(), big.NewInt(2470038), big.NewInt(5))
	if tx != nil {
		//a:=tx.
		//u(tx.Data.Recipient, tx.From, tx.Data.Amount, tx.Data.Hash)
		//froms = append(froms, tx.From)
		//tos = append(tos, tx.Data.Recipient)
		//blocks = append(blocks, tx)
		//logs.Infof("Tx Detail: %s", tx)

		//		log.Info(fmt.Sprintf("Tx Detail: %s", tx.Data.Hash))
		//log2.New()
		Amount,_:=hexutil.DecodeBig(tx.Amount)
		BlockNumber,_:=hexutil.DecodeBig(tx.BlockNumber)
		fmt.Println("tx From:", tx.From.String())
		fmt.Println("tx BlockNumber:", BlockNumber)
		fmt.Println("tx AccountNonce:", tx.AccountNonce)
		fmt.Println("tx Amount Ex:", Amount)
		fmt.Println("tx Recipient:", tx.Recipient.String())
	} else {
		//		mx.log.Info("Tx Detail is null")
	}

}
