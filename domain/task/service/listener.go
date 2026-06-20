package service

import (
	"context"
	"log"

	"github.com/Y1le/godolist/domain/event"
	userevent "github.com/Y1le/godolist/domain/user/event"
)

func (t *TaskDomainImpl) OnUserRenamed(ctx context.Context, e event.Event) error {
	renamed, ok := e.(*userevent.UserRenamed)
	if !ok {
		return nil
	}
	if err := t.repo.BumpUserName(ctx, nil, renamed.UserID, renamed.NewUsername); err != nil {
		log.Printf("[task] BumpUserName failed: %v", err)
		return err
	}
	return nil
}