package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/Y1le/godolist/domain/task/entity"
	"github.com/Y1le/godolist/interfaces/types"
)

type Task interface {
	TaskBase
	TaskQuery
}

type TaskBase interface {
	CreateTask(ctx context.Context, tx *gorm.DB, task *entity.Task) (*entity.Task, error)
	UpdateTask(ctx context.Context, tx *gorm.DB, task *entity.Task) error
	ListTaskByUid(ctx context.Context, tx *gorm.DB, uid uint, p types.Pagination) ([]*entity.Task, int64, error)
	FindTaskByTid(ctx context.Context, tx *gorm.DB, tid, uid uint) (*entity.Task, error)
	SearchTask(ctx context.Context, tx *gorm.DB, uid uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error)
	DeleteTask(ctx context.Context, tx *gorm.DB, uid, tid uint) error

	// BumpUserName is invoked by the listener that handles user.renamed.
	// Implementation updates task rows that hold a redundant user_name copy.
	BumpUserName(ctx context.Context, tx *gorm.DB, uid uint, newName string) error
}

type TaskQuery interface {
	// ...
}