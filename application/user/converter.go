package user

import (
	"github.com/Y1le/godolist/domain/user/entity"
	"github.com/Y1le/godolist/interfaces/types"
)

func LoginResponse(u *entity.User, token string) *types.TokenData {
	return &types.TokenData{
		User: types.UserResp{
			ID:       u.ID,
			UserName: u.Username,
			CreateAt: u.CreatedAt.Unix(),
		},
		Token: token,
	}
}

func RegisterResponse(u *entity.User) *types.UserResp {
	return &types.UserResp{
		ID:       u.ID,
		UserName: u.Username,
		CreateAt: u.CreatedAt.Unix(),
	}
}
