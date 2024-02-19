package job

import (
	"github.com/google/wire"
)

var ProviderCronjobSet = wire.NewSet(
	NewGoNewVersion,
)
