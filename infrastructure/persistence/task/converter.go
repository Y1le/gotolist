package task

import (
	"github.com/Y1le/gotolist/domain/task/entity"
)

func Entity2PO(task *entity.Task) *Task {
	return &Task{
		Uid:       task.Uid,
		UserName:  task.UserName,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

// PO2Entity no longer takes a *user.User. user_name is read from the
// task row itself; the task BC keeps its own copy and refreshes it
// via the user.renamed event listener.
func PO2Entity(t *Task) *entity.Task {
	return &entity.Task{
		Id:        t.ID,
		Uid:       t.Uid,
		UserName:  t.UserName,
		Title:     t.Title,
		Status:    t.Status,
		Content:   t.Content,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
	}
}