package stellar

import (
	"context"

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/clients/stellarcore"

	proto "github.com/stellar/go/protocols/stellarcore"
	// for building transactions
	//"github.com/stellar/go/build"
)

// client defines typed wrappers for the Steller API.
type client struct {
	*horizon.Client
	core *stellarcore.Client // allow talking directly to core
}

// Dial just associates URLs with the client, it does not actually try to connect (yet)
func Dial(horizonURL string, coreURL string) (Client, error) {
	// TODO: try to do basic connction to both URLs to detect URL problems at Dial time
	// TODO: set up HTTP?
	return &client{
		Client: &horizon.Client{URL: horizonURL},
		core:   &stellarcore.Client{URL: coreURL},
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
