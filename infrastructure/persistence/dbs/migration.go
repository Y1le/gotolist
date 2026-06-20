package dbs

import (
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/outbox"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/task"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/user"
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