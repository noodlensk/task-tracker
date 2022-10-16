// Package publisher provides primitives to interact with the asyncapi
//
// Code generated by https://github.com/asyncapi/generator/ DO NOT EDIT.
package publisher

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

func (c PublisherClient) UserCreated(ctx context.Context, e UserCreated) error {
	data, err := json.Marshal(e)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}

	return c.publisher.Publish("users-cud.created", message.NewMessage(watermill.NewShortUUID(), data))
}

func (c PublisherClient) UserUpdated(ctx context.Context, e UserUpdated) error {
	data, err := json.Marshal(e)
	if err != nil {
		return errors.Wrap(err, "marshal data")
	}

	return c.publisher.Publish("users-cud.updated", message.NewMessage(watermill.NewShortUUID(), data))
}
