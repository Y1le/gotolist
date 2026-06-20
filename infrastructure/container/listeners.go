package container

import (
	"github.com/Y1le/gotolist/domain/event"
	"github.com/Y1le/gotolist/domain/task/service"
	"github.com/Y1le/gotolist/infrastructure/eventbus"
)

func registerListeners(bus *eventbus.InProcBus, taskDomain service.TaskDomain) {
	bus.Subscribe("user.renamed", event.HandlerFunc(taskDomain.OnUserRenamed))
}