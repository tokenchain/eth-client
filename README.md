Node Wrapper Client SDK
===============

[![License: LGPL v3](https://img.shields.io/badge/License-LGPL%20v3-blue.svg)](https://www.gnu.org/licenses/lgpl-3.0)
[![Travis](https://img.shields.io/travis/rust-lang/rust.svg)](https://travis-ci.org/getamis/eth-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/getamis/eth-client)](https://goreportcard.com/report/github.com/getamis/eth-client)


A Golang client library to communicate with Ethereum and Stellar RPC servers.
* Implements most of JSON-RPC methods and several client-specific methods.
* Provides a high-level interface to **propose/get validators** on an Istanbul blockchain.
* Provides a high-level interface to **create private contracts** on a Quorum blockchain.
* Provides a bare bones wrapper to **send simple transactions** on a Stellar blockchain.
* Provides a generic client interface for all Dial()ed endpoints:
* `GetInfo()`
* `SendAmount(from, to, amount)`
* `GetBalance(account)`


```

go get -u github.com/ethereum/go-ethereum
go get -u github.com/getamis/eth-client/
go get -u github.com/stellar/go
go get -u github.com/inconshreveable/log15

```

Usage
-----

```golang
package main

import (
	"context"
	"fmt"

	// Generic interface common to all endpoint types
	//"github.com/Blockdaemon/node-client-sdk/client"

	// Ethereum specific interface
	"github.com/Blockdaemon/node-client-sdk/eth"

	// Stellar/Horzion specific interface
	//"github.com/Blockdaemon/node-client-sdk/stellar"
)

func main() {
	// Note that by default, HTTP does not support the mining API
	// If mining is not enabled on the node's HTTP endpoint, replace
	// with local IPC socket filename
	url := "http://127.0.0.1:8545"
	client, err := eth.Dial(url)

	// For Stellar, something like:
	//url := "http://127.0.0.1:8000"
	//coreurl := "http://127.0.0.1:11626"
	//passphrase := ""
	//client, err := stellar.Dial(url, coreurl, passphrase)

	if err != nil {
		fmt.Println("Failed to dial, url: ", url, ", err: ", err)
		return
	}

	// Should work for all types of Dial()ed endpoints
	balance, err := client.GetBalance(context.Background(),
		"de155b6f2aead0474c7428424dec755170e97f76")
	if err != nil {
		fmt.Println("Failed to get balance, err: ", err)
		return
	}
	fmt.Println("balance: ", balance)

	// Ethereum specific
	err = client.SetMiningAccount(context.Background(),
		"de155b6f2aead0474c7428424dec755170e97f76")
	if err != nil {
		fmt.Println("Failed to set etherbase, err: ", err)
		return
	}

	err = client.StartMining(context.Background())
	if err != nil {
		fmt.Println("Failed to start mining, err: ", err)
		return
	}
	fmt.Println("start mining")
}

```

Implemented JSON-RPC methods
----------------------------

## Ethereum:

* admin_addPeer
* admin_adminPeers
* admin_nodeInfo
* eth_blockNumber
* eth_sendRawTransaction
* eth_getBlockByHash
* eth_getBlockByNumber
* eth_getBlockByHash
* eth_getBlockByNumber
* eth_getTransactionByHash
* eth_getBlockTransactionCountByHash
* eth_getTransactionByBlockHashAndIndex
* eth_getTransactionReceipt
* eth_syncing
* eth_getBalance
* eth_getStorageAt
* eth_getCode
* eth_getBlockTransactionCountByNumber
* eth_call
* eth_gasPrice
* eth_estimateGas
* eth_sendRawTransaction
* miner_setEtherbase
* miner_startMining
* miner_stopMining
* net_version
* supportedModules
* logs
* newHeads
* eth_getLogs

### Istanbul-only JSON-RPC methods
To use these methods, make sure that
* Server is running on [Istanbul consensus](https://github.com/ethereum/EIPs/issues/650).
* Connect to server through `istanbul.Dial` function (not the original [Geth client](https://github.com/ethereum/go-ethereum/tree/master/ethclient)).

Methods:

* istanbul_getValidators
* istanbul_propose

### Quorum-only JSON-RPC methods

To use these methods, make sure that
* Server is running on [Quorum blockchain](https://github.com/jpmorganchase/quorum/wiki)
* Connect to server through `quorum.Dial` function (not the original [Geth client](https://github.com/ethereum/go-ethereum/tree/master/ethclient)).

Methods:

* quorum_privateContract
* quorum_contract

## Stellar:

To use these methods, make sure that
* Server is running a [Stellar blockchain](https://github.com/stellar/packages)
* Make sure the horizon service is running as well.
* Connect to server through `stellar.Dial` function. You will need to supply both the core and horizon endpoints, as well as the network passphrase

Methods:

* core.Info
* horizon.SubmitTransaction
* horizon.LoadAccount

Contributing
------------

Feel free to contribute to this repository.

1. Fork it!
2. Create your feature branch: git checkout -b my-new-feature
3. Commit your changes: git commit -am 'Add some feature'
4. Push to the branch: git push origin my-new-feature
5. Submit a pull request

Reference
---------

* https://github.com/ethereum/go-ethereum
* https://github.com/ethereum/wiki/wiki/JSON-RPC
* https://github.com/ethereum/EIPs/issues/650
* https://github.com/jpmorganchase/quorum
* https://github.com/getamis/istanbul-tools
* https://github.com/stellar/go/clients/stellarcore
* https://github.com/stellar/go/clients/horizon
* https://github.com/stellar/go/build
