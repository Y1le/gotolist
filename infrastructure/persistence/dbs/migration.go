package dbs

import (
	"github.com/Y1le/gotolist/infrastructure/persistence/outbox"
	"github.com/Y1le/gotolist/infrastructure/persistence/task"
	"github.com/Y1le/gotolist/infrastructure/persistence/user"
)

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&user.User{},
			&task.Task{},
			&outbox.OutboxPO{},
		)
	if err != nil {
		return
	}
}