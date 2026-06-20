package container

import (
	"github.com/CocaineCong/todolist-ddd/domain/event"
	taskSrv "github.com/CocaineCong/todolist-ddd/domain/task/service"
	userevent "github.com/CocaineCong/todolist-ddd/domain/user/event"
	userSrv "github.com/CocaineCong/todolist-ddd/domain/user/service"
	"github.com/CocaineCong/todolist-ddd/infrastructure/auth"
	"github.com/CocaineCong/todolist-ddd/infrastructure/encrypt"
	"github.com/CocaineCong/todolist-ddd/infrastructure/eventbus"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/dbs"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/outbox"

	taskApp "github.com/CocaineCong/todolist-ddd/application/task"
	userApp "github.com/CocaineCong/todolist-ddd/application/user"
)

func LoadingDomain(bus *eventbus.InProcBus) {
	repos := persistence.NewRepositories(dbs.DB)
	jwtService := auth.NewJWTTokenService()
	pwdEncryptService := encrypt.NewPwdEncryptService()
	eventStore := outbox.New()

	bus.Register("user.created", func() event.Event { return &userevent.UserCreated{} })
	bus.Register("user.renamed", func() event.Event { return &userevent.UserRenamed{} })

	userDomain := userSrv.NewUserDomainImpl(dbs.DB, repos.User, pwdEncryptService, eventStore)
	userApp.GetServiceImpl(userDomain, jwtService)

	taskDomain := taskSrv.NewTaskDomainImpl(dbs.DB, repos.Task, eventStore)
	taskApp.GetServiceImpl(taskDomain)

	registerListeners(bus, taskDomain)
}
