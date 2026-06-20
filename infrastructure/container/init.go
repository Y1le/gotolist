package container

import (
	"github.com/Y1le/gotolist/domain/event"
	taskSrv "github.com/Y1le/gotolist/domain/task/service"
	userevent "github.com/Y1le/gotolist/domain/user/event"
	userSrv "github.com/Y1le/gotolist/domain/user/service"
	"github.com/Y1le/gotolist/infrastructure/auth"
	"github.com/Y1le/gotolist/infrastructure/encrypt"
	"github.com/Y1le/gotolist/infrastructure/eventbus"
	"github.com/Y1le/gotolist/infrastructure/persistence"
	"github.com/Y1le/gotolist/infrastructure/persistence/dbs"
	"github.com/Y1le/gotolist/infrastructure/persistence/outbox"

	taskApp "github.com/Y1le/gotolist/application/task"
	userApp "github.com/Y1le/gotolist/application/user"
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
