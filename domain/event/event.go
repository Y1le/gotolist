package event

import (
	"context"
	"time"
)

type Event interface {
	EventName() string
	OccurredAt() time.Time
}

type Handler interface {
	Handle(ctx context.Context, e Event) error
}

type HandlerFunc func(ctx context.Context, e Event) error

func (f HandlerFunc) Handle(ctx context.Context, e Event) error { return f(ctx, e) }

type Bus interface {
	Publish(ctx context.Context, e Event) error
	Subscribe(eventName string, h Handler)
	Start(ctx context.Context) error
}