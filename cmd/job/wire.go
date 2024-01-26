//go:build wireinject
// +build wireinject

package job

import (
	"github.com/google/wire"
	"wlb/internal/clients"
	"wlb/internal/conf"
	"wlb/internal/service/job"
)

var ProviderSet = wire.NewSet(
	// ProviderCronjobSet,
	job.ProviderCronjobSet,
	// 基础层
	clients.ProviderClientsSet,
)

func NewGoNewVersionJob(data *conf.Data) (*job.GoNewVersionJob, func(), error) {
	panic(wire.Build(
		ProviderSet,
	))
}
