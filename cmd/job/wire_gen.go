// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package job

import (
	"github.com/google/wire"
	"wlb/internal/clients"
	"wlb/internal/conf"
	"wlb/internal/service/job"
)

// Injectors from wire.go:

func NewGoNewVersionJob(data *conf.Data) (*job.GoNewVersionJob, func(), error) {
	client := clients.NewLarkClient(data)
	goNewVersionJob := job.NewGoNewVersionJob(client)
	return goNewVersionJob, func() {
	}, nil
}

// wire.go:

var ProviderSet = wire.NewSet(job.ProviderCronjobSet, clients.ProviderClientsSet)