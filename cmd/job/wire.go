//go:build wireinject
// +build wireinject

package job

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"
	"wlb/internal/biz"
	"wlb/internal/clients"
	"wlb/internal/conf"
	"wlb/internal/data"
	"wlb/internal/service/job"
)

var ProviderSet = wire.NewSet(
	// ProviderCronjobSet,
	job.ProviderCronjobSet,
	biz.BizProviderSet,
	// 基础层
	clients.ProviderClientsSet,
	data.DataProviderSet,
)

func NewGoNewVersionJob(data *conf.Data, logger log.Logger) (*job.GoNewVersion, func(), error) {
	panic(wire.Build(
		ProviderSet,
	))
}
