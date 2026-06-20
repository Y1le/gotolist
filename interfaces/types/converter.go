package types

import (
	"github.com/Y1le/godolist/domain/task/entity"
	ue "github.com/Y1le/godolist/domain/user/entity"
)

func Entity2TaskResp(task *entity.Task) *TaskResp {
	return &TaskResp{
		ID:        task.Id,
		Title:     task.Title,
		Content:   task.Content,
		View:      0,
		Status:    task.Status,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func UserReq2Entity(user *UserReq) *ue.User {
	return &ue.User{
		Username: user.UserName,
		Password: user.Password,
	}
}
