// Package adapters provides primitives to interact with the asyncapi
//
// Code generated by https://github.com/asyncapi/generator/ DO NOT EDIT.
package adapters

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
)

func NewPublisherClient(pub message.Publisher) *PublisherClient {
	return &PublisherClient{
		publisher: pub,
	}
}

type PublisherClient struct {
	publisher message.Publisher
}

func (c PublisherClient) TaskEstimated(ctx context.Context, e TaskEstimated) error {
	data, err := json.Marshal(e)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}

	return c.publisher.Publish("accounting.task_estimated", message.NewMessage(watermill.NewShortUUID(), data))
}
