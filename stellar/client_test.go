package stellar

// Verfiy that client implements the Client interface.
var (
	_ = Client(&client{})
)
