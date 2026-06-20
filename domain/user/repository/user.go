package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/Y1le/godolist/domain/user/entity"
)

type User interface {
	UserBase
}

// UserBase uses *gorm.DB so domain service can wrap the calls in a
// transaction and atomically append a domain event to event_outbox.
type UserBase interface {
	CreateUser(ctx context.Context, tx *gorm.DB, user *entity.User) (*entity.User, error)
	GetUserByName(ctx context.Context, tx *gorm.DB, name string) (*entity.User, error)
	GetUserByID(ctx context.Context, tx *gorm.DB, id uint) (*entity.User, error)
}