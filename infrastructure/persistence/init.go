package persistence

import (
	"gorm.io/gorm"

	tRepo "github.com/Y1le/godolist/domain/task/repository"
	uRepo "github.com/Y1le/godolist/domain/user/repository"
	"github.com/Y1le/godolist/infrastructure/persistence/task"
	"github.com/Y1le/godolist/infrastructure/persistence/user"
)

type Repositories struct {
	User uRepo.User
	Task tRepo.Task
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewRepository(db),
		Task: task.NewRepository(db),
	}
}
