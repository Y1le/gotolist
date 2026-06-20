package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/Y1le/godolist/domain/user/entity"
	"github.com/Y1le/godolist/domain/user/repository"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.User {
	return &RepositoryImpl{db: db}
}

// withDB picks a connection: a passed-in tx if non-nil, else the global pool.
func (r *RepositoryImpl) withDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *RepositoryImpl) CreateUser(ctx context.Context, tx *gorm.DB, user *entity.User) (*entity.User, error) {
	u := Entity2PO(user)
	err := r.withDB(tx).WithContext(ctx).Model(&User{}).Create(u).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(u), nil
}

func (r *RepositoryImpl) GetUserByName(ctx context.Context, tx *gorm.DB, username string) (*entity.User, error) {
	var u *User
	err := r.withDB(tx).WithContext(ctx).Model(&User{}).Where("user_name = ?", username).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(u), nil
}

func (r *RepositoryImpl) GetUserByID(ctx context.Context, tx *gorm.DB, id uint) (*entity.User, error) {
	var u *User
	err := r.withDB(tx).WithContext(ctx).Model(&User{}).Where("id = ?", id).Find(&u).Error
	if err != nil {
		return nil, err
	}
	if u.ID == 0 {
		return nil, errors.New("user not found")
	}
	return PO2Entity(u), nil
}