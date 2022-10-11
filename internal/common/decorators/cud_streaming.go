package decorator

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type commandCUDStreamingDecorator[C any, R any] struct {
	base      CommandHandler[C, R]
	publisher message.Publisher
	topic     string
}

func (d commandCUDStreamingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	res, err := d.base.Handle(ctx, cmd)
	if err != nil {
		return res, err
	}

	if d.publisher == nil { // TODO: make it more smart for tests
		return res, err
	}

	data, err := json.Marshal(res)
	if err != nil {
		return res, err
	}

	if err := d.publisher.Publish(d.topic, message.NewMessage(watermill.NewShortUUID(), data)); err != nil {
		return res, err
	}

	return res, err
}
