package stellar

import (
	"context"

	generic "github.com/Blockdaemon/node-client-sdk/client"

	proto "github.com/stellar/go/protocols/stellarcore"
	// for building transactions
	//"github.com/stellar/go/build"
	// horizon
	//"github.com/stellar/go/clients/horizon"
)

type Client interface {
	generic.Client

	// horizon
	SubmitTransaction(ctx context.Context, envelope string) (resp *proto.TXResponse, err error) // simple wrapper to existing SubmitTransaction
}
