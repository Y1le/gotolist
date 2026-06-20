package event

import "time"

type UserCreated struct {
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func (UserCreated) EventName() string       { return "user.created" }
func (e UserCreated) OccurredAt() time.Time { return e.CreatedAt }

type UserRenamed struct {
	UserID      uint      `json:"user_id"`
	OldUsername string    `json:"old_username"`
	NewUsername string    `json:"new_username"`
	RenamedAt   time.Time `json:"renamed_at"`
}

func (UserRenamed) EventName() string       { return "user.renamed" }
func (e UserRenamed) OccurredAt() time.Time { return e.RenamedAt }