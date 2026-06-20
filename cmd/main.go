package main

import (
	"context"

	config "github.com/Y1le/godolist/conf"
	"github.com/Y1le/godolist/infrastructure/common/log"
	"github.com/Y1le/godolist/infrastructure/container"
	"github.com/Y1le/godolist/infrastructure/eventbus"
	"github.com/Y1le/godolist/infrastructure/persistence/dbs"
	"github.com/Y1le/godolist/interfaces/adapter/initialize"
)

func main() {
	loadingInfra()

	bus := eventbus.NewInProcBus(dbs.DB)
	container.LoadingDomain(bus)
	go bus.Start(context.Background())

	r := initialize.NewRouter()
	addr := ":3001"
	if config.Conf != nil && config.Conf.Server != nil && config.Conf.Server.Port != "" {
		addr = config.Conf.Server.Port
	}
	_ = r.Run(addr)
}

func loadingInfra() {
	config.InitConfig()
	log.InitLog()
	dbs.MySQLInit()
}