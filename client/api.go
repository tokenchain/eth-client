package client

import (
	"context"
)

type Client interface {
	Close()

	GetInfo(ctx context.Context) (string, error)
	SendAmount(ctx context.Context, fromPriv, toPub, amount string) error
	GetBalance(ctx context.Context, account string) (string, error)
}
