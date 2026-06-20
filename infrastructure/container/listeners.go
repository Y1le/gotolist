package container

import (
	"github.com/CocaineCong/todolist-ddd/domain/event"
	"github.com/CocaineCong/todolist-ddd/domain/task/service"
	"github.com/CocaineCong/todolist-ddd/infrastructure/eventbus"
)

func registerListeners(bus *eventbus.InProcBus, taskDomain service.TaskDomain) {
	bus.Subscribe("user.renamed", event.HandlerFunc(taskDomain.OnUserRenamed))
}