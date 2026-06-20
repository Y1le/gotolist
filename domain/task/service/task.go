package service

import (
	"context"

	"gorm.io/gorm"

	"github.com/Y1le/godolist/domain/event"
	"github.com/Y1le/godolist/domain/task/entity"
	"github.com/Y1le/godolist/domain/task/repository"
	"github.com/Y1le/godolist/infrastructure/persistence/outbox"
	"github.com/Y1le/godolist/interfaces/types"
)

type UserContextProvider interface {
	GetCurrentUserID(ctx context.Context) (uint, error)
	GetCurrentUserName(ctx context.Context) (string, error)
}

type TaskDomain interface {
	CreateTask(ctx context.Context, in *entity.Task) (*entity.Task, error)
	FindTaskByTid(ctx context.Context, taskId, userId uint) (*entity.Task, error)
	ListTaskByUid(ctx context.Context, userId uint, p types.Pagination) ([]*entity.Task, int64, error)
	UpdateTask(ctx context.Context, in *entity.Task) error
	SearchTask(ctx context.Context, userId uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error)
	DeleteTask(ctx context.Context, uid, tid uint) error
	OnUserRenamed(ctx context.Context, e event.Event) error
}

type TaskDomainImpl struct {
	db    *gorm.DB
	repo  repository.Task
	store *outbox.Outbox
}

func NewTaskDomainImpl(db *gorm.DB, repo repository.Task, store *outbox.Outbox) TaskDomain {
	return &TaskDomainImpl{db: db, repo: repo, store: store}
}

func (t *TaskDomainImpl) CreateTask(ctx context.Context, in *entity.Task) (*entity.Task, error) {
	return t.repo.CreateTask(ctx, nil, in)
}

func (t *TaskDomainImpl) FindTaskByTid(ctx context.Context, taskId, userId uint) (*entity.Task, error) {
	return t.repo.FindTaskByTid(ctx, nil, taskId, userId)
}

func (t *TaskDomainImpl) ListTaskByUid(ctx context.Context, userId uint, p types.Pagination) ([]*entity.Task, int64, error) {
	return t.repo.ListTaskByUid(ctx, nil, userId, p)
}

func (t *TaskDomainImpl) UpdateTask(ctx context.Context, task *entity.Task) error {
	return t.repo.UpdateTask(ctx, nil, task)
}

func (t *TaskDomainImpl) SearchTask(ctx context.Context, userId uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error) {
	return t.repo.SearchTask(ctx, nil, userId, keyword, p)
}

func (t *TaskDomainImpl) DeleteTask(ctx context.Context, uid, tid uint) error {
	return t.repo.DeleteTask(ctx, nil, uid, tid)
}
