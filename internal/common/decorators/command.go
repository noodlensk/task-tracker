package decorator

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

func ApplyCommandDecorators[H any, R any](handler CommandHandler[H, R], publisher message.Publisher, topic string) CommandHandler[H, R] {
	return commandCUDStreamingDecorator[H, R]{
		base: commandCUDStreamingDecorator[H, R]{
			base:      handler,
			publisher: publisher,
			topic:     topic,
		},
	}
}

type CommandHandler[C any, R any] interface {
	Handle(ctx context.Context, c C) (R, error)
}
