package main

import (
	"context"

	config "github.com/CocaineCong/todolist-ddd/conf"
	"github.com/CocaineCong/todolist-ddd/infrastructure/common/log"
	"github.com/CocaineCong/todolist-ddd/infrastructure/container"
	"github.com/CocaineCong/todolist-ddd/infrastructure/eventbus"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/dbs"
	"github.com/CocaineCong/todolist-ddd/interfaces/adapter/initialize"
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