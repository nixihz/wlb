package data

import (
	"wlb/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewGoVersionRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{}, cleanup, nil
}
