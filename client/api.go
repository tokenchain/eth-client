package client

import (
	"context"
)

type Client interface {
	Close()

	GetInfo(ctx context.Context) (string, error)
	GenerateKey(ctx context.Context) (address, private string, err error)
	SendAmount(ctx context.Context, fromPriv, toPub, amount string) error
	GetBalance(ctx context.Context, account string) (string, error)
}
