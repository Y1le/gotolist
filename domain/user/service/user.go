package service

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	userevent "github.com/Y1le/godolist/domain/user/event"
	"github.com/Y1le/godolist/domain/user/entity"
	"github.com/Y1le/godolist/domain/user/repository"
	"github.com/Y1le/godolist/infrastructure/persistence/outbox"
)

type UserDomain interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserByName(ctx context.Context, name string) (*entity.User, error)
	GetUserDetail(ctx context.Context, id uint) (*entity.User, error)
	CheckUserPwd(ctx context.Context, user *entity.User, src string) error
	RenameUser(ctx context.Context, id uint, newName string) error
}

type UserDomainImpl struct {
	db      *gorm.DB
	repo    repository.User
	encrypt repository.PwdEncrypt
	store   *outbox.Outbox
}

func NewUserDomainImpl(
	db *gorm.DB,
	repo repository.User,
	encrypt repository.PwdEncrypt,
	store *outbox.Outbox,
) UserDomain {
	return &UserDomainImpl{db: db, repo: repo, encrypt: encrypt, store: store}
}

func (u *UserDomainImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	encryptPwd, err := u.encrypt.Encrypt([]byte(user.Password))
	if err != nil {
		return nil, err
	}
	if err := user.SetPwd(encryptPwd); err != nil {
		return nil, err
	}

	var created *entity.User
	err = u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		out, err := u.repo.CreateUser(ctx, tx, user)
		if err != nil {
			return err
		}
		created = out
		return u.store.Append(ctx, tx, userevent.UserCreated{
			UserID:    out.ID,
			Username:  out.Username,
			CreatedAt: out.CreatedAt,
		})
	})
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (u *UserDomainImpl) FindUserByName(ctx context.Context, username string) (*entity.User, error) {
	return u.repo.GetUserByName(ctx, nil, username)
}

func (u *UserDomainImpl) GetUserDetail(ctx context.Context, id uint) (*entity.User, error) {
	return u.repo.GetUserByID(ctx, nil, id)
}

func (u *UserDomainImpl) CheckUserPwd(ctx context.Context, user *entity.User, src string) error {
	if u.encrypt.Check([]byte(user.Password), []byte(src)) {
		return nil
	}
	return errors.New("wrong password")
}

func (u *UserDomainImpl) RenameUser(ctx context.Context, id uint, newName string) error {
	return u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		cur, err := u.repo.GetUserByID(ctx, tx, id)
		if err != nil {
			return err
		}
		if cur.Username == newName {
			return nil
		}
		oldName := cur.Username
		if err := tx.WithContext(ctx).
			Exec("UPDATE user SET user_name = ?, updated_at = ? WHERE id = ?",
				newName, time.Now().Unix(), id).Error; err != nil {
			return err
		}
		return u.store.Append(ctx, tx, userevent.UserRenamed{
			UserID:      id,
			OldUsername: oldName,
			NewUsername: newName,
			RenamedAt:   time.Now(),
		})
	})
}