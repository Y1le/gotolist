package container

import (
	"github.com/Y1le/godolist/domain/event"
	taskSrv "github.com/Y1le/godolist/domain/task/service"
	userevent "github.com/Y1le/godolist/domain/user/event"
	userSrv "github.com/Y1le/godolist/domain/user/service"
	"github.com/Y1le/godolist/infrastructure/auth"
	"github.com/Y1le/godolist/infrastructure/encrypt"
	"github.com/Y1le/godolist/infrastructure/eventbus"
	"github.com/Y1le/godolist/infrastructure/persistence"
	"github.com/Y1le/godolist/infrastructure/persistence/dbs"
	"github.com/Y1le/godolist/infrastructure/persistence/outbox"

	taskApp "github.com/Y1le/godolist/application/task"
	userApp "github.com/Y1le/godolist/application/user"
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
