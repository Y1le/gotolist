package container

import (
	"github.com/Y1le/godolist/domain/event"
	"github.com/Y1le/godolist/domain/task/service"
	"github.com/Y1le/godolist/infrastructure/eventbus"
)

func registerListeners(bus *eventbus.InProcBus, taskDomain service.TaskDomain) {
	bus.Subscribe("user.renamed", event.HandlerFunc(taskDomain.OnUserRenamed))
}