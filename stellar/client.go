package stellar

import (
	"context"
	"encoding/json"

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/clients/stellarcore"

	proto "github.com/stellar/go/protocols/stellarcore"
	// for building transactions
	"github.com/stellar/go/build"
)

// client defines typed wrappers for the Steller API.
type client struct {
	*horizon.Client
	core       *stellarcore.Client // allow talking directly to core
	passphrase string              // network passphrase
}

// Dial just associates URLs with the client, it does not actually try to connect (yet)
func Dial(horizonURL string, coreURL string, passphrase string) (Client, error) {
	// TODO: try to do basic connction to both URLs to detect
	// URL/passphrase problems at Dial time
	// TODO: set up HTTP?
	return &client{
		Client:     &horizon.Client{URL: horizonURL},
		core:       &stellarcore.Client{URL: coreURL},
		passphrase: passphrase,
	}, nil
}

func (c *client) Close() {
	// TODO: close HTTP if set up in Dial?
}

func (c *client) Info(ctx context.Context) (resp *proto.InfoResponse, err error) {
	return c.core.Info(ctx)
}

func (c *client) SubmitTransaction(ctx context.Context, envelope string) (resp *proto.TXResponse, err error) {
	return c.core.SubmitTransaction(ctx, envelope)
}

func (c *client) BuildAndSendTransaction(ctx context.Context, from, to, amount string) (resp *proto.TXResponse, err error) {
	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: from},
		build.Network{Passphrase: c.passphrase},
		build.AutoSequence{SequenceProvider: c},
		build.Payment(
			build.Destination{AddressOrSeed: to},
			build.NativeAmount{Amount: amount},
		),
	)
	if err != nil {
		return nil, err
	}

	txe, err := tx.Sign(from)
	if err != nil {
		return nil, err
	}

	txeB64, err := txe.Base64()
	if err != nil {
		return nil, err
	}

	return c.SubmitTransaction(ctx, txeB64)
}

func (c *client) GetBalance(ctx context.Context, address string) (resp string, err error) {
	account, err := c.LoadAccount(address)
	if err != nil {
		return "", err
	}

	balances, err := json.Marshal(account.Balances)
	if err != nil {
		return "", err
	}

	return string(balances), err
}
