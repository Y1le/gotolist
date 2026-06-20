package outbox

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/CocaineCong/todolist-ddd/domain/event"
)

type OutboxPO struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	EventName   string `gorm:"size:64;not null;index"`
	Payload     string `gorm:"type:longtext;not null"`
	Status      string `gorm:"size:16;not null;default:'pending';index"`
	Attempts    int    `gorm:"not null;default:0"`
	LastError   string `gorm:"type:text"`
	CreatedAt   int64  `gorm:"not null;autoCreateTime"`
	ProcessedAt int64  `gorm:"not null;default:0"`
}

func (OutboxPO) TableName() string { return "event_outbox" }

type Outbox struct{}

func New() *Outbox { return &Outbox{} }

var ErrNilTx = &outboxError{"Append: tx is nil"}

func (o *Outbox) Append(ctx context.Context, tx *gorm.DB, e event.Event) error {
	if tx == nil {
		return ErrNilTx
	}
	payload, err := json.Marshal(e)
	if err != nil {
		return err
	}
	row := &OutboxPO{
		EventName: e.EventName(),
		Payload:   string(payload),
		Status:    "pending",
	}
	return tx.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(row).Error
}

type outboxError struct{ s string }

func (e *outboxError) Error() string { return e.s }