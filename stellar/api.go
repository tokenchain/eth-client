package stellar

import (
	"context"

	proto "github.com/stellar/go/protocols/stellarcore"
	// for building transactions
	//"github.com/stellar/go/build"
	// horizon
	//"github.com/stellar/go/clients/horizon"
)

type Client interface {
	Close()

	// stellar-core
	Info(ctx context.Context) (resp *proto.InfoResponse, err error)

	// horizon
	SubmitTransaction(ctx context.Context, envelope string) (resp *proto.TXResponse, err error)               // simple wrapper to existing SubmitTransaction
	BuildAndSendTransaction(ctx context.Context, from, to, amount string) (resp *proto.TXResponse, err error) // build, sign, and submit
	GetBalance(ctx context.Context, address string) (resp string, err error)
}
