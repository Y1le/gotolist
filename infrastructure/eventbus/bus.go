package eventbus

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/Y1le/gotolist/domain/event"
	"github.com/Y1le/gotolist/infrastructure/persistence/outbox"
)

type InProcBus struct {
	mu       sync.RWMutex
	subs     map[string][]event.Handler
	factories map[string]func() event.Event
	outboxDB *gorm.DB
}

func NewInProcBus(db *gorm.DB) *InProcBus {
	return &InProcBus{
		subs:      make(map[string][]event.Handler),
		factories: make(map[string]func() event.Event),
		outboxDB:  db,
	}
}

// Register wires an event name to both:
//   - a typed factory used to decode outbox rows back into a concrete event
//   - one or more subscribers (call Subscribe after Register)
func (b *InProcBus) Register(name string, factory func() event.Event) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.factories[name] = factory
}

func (b *InProcBus) Subscribe(eventName string, h event.Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subs[eventName] = append(b.subs[eventName], h)
}

func (b *InProcBus) Publish(ctx context.Context, e event.Event) error {
	b.mu.RLock()
	hs := append([]event.Handler(nil), b.subs[e.EventName()]...)
	b.mu.RUnlock()
	for _, h := range hs {
		if err := h.Handle(ctx, e); err != nil {
			log.Printf("[eventbus] handler for %s failed: %v", e.EventName(), err)
		}
	}
	return nil
}

func (b *InProcBus) Start(ctx context.Context) error {
	tick := time.NewTicker(500 * time.Millisecond)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-tick.C:
			b.drainOnce(ctx)
		}
	}
}

func (b *InProcBus) drainOnce(ctx context.Context) {
	var rows []outbox.OutboxPO
	if err := b.outboxDB.WithContext(ctx).
		Where("status = ?", "pending").
		Order("id ASC").
		Limit(100).
		Find(&rows).Error; err != nil {
		log.Printf("[eventbus] outbox scan error: %v", err)
		return
	}
	for _, row := range rows {
		e, err := b.decodeRow(row)
		if err != nil {
			b.markFailed(ctx, row.ID, err)
			continue
		}
		_ = b.Publish(ctx, e)
		b.outboxDB.WithContext(ctx).Model(&outbox.OutboxPO{}).
			Where("id = ?", row.ID).
			Updates(map[string]interface{}{
				"status":       "done",
				"processed_at": time.Now().Unix(),
			})
	}
}

func (b *InProcBus) markFailed(ctx context.Context, id uint64, err error) {
	b.outboxDB.WithContext(ctx).Model(&outbox.OutboxPO{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     "failed",
			"attempts":   gorm.Expr("attempts + 1"),
			"last_error": err.Error(),
		})
}

func (b *InProcBus) decodeRow(row outbox.OutboxPO) (event.Event, error) {
	b.mu.RLock()
	factory, ok := b.factories[row.EventName]
	b.mu.RUnlock()
	if !ok {
		return nil, errors.New("no factory registered for " + row.EventName)
	}
	e := factory()
	if err := json.Unmarshal([]byte(row.Payload), e); err != nil {
		return nil, err
	}
	return e, nil
}