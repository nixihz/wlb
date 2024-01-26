package clients

import (
	"github.com/google/wire"
)

// ProviderClientsSet is biz providers.
var ProviderClientsSet = wire.NewSet(
	NewLarkClient,
)
