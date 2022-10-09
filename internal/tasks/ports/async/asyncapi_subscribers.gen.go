// Package async provides primitives to interact with the asyncapi
//
// Code generated by https://github.com/asyncapi/generator/ DO NOT EDIT.
package async

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
)

func Register(application ServerInterface, router *message.Router, subscriber message.Subscriber) error {
	asyncServer := ServerInterfaceWrapper{app: application}

	router.AddNoPublisherHandler(
		"tasks.reassign",
		"tasks.reassign",
		subscriber,
		asyncServer.TasksReAssign,
	)

	router.AddNoPublisherHandler(
		"users.created",
		"users.created",
		subscriber,
		asyncServer.UserCreated,
	)

	router.AddNoPublisherHandler(
		"users.updated",
		"users.updated",
		subscriber,
		asyncServer.UserUpdated,
	)

	return nil
}

type ServerInterface interface {
	TasksReAssign(ctx context.Context, e TasksReAssign) error

	UserCreated(ctx context.Context, e UserCreated) error

	UserUpdated(ctx context.Context, e UserUpdated) error
}

type ServerInterfaceWrapper struct {
	app ServerInterface
}

func (s ServerInterfaceWrapper) TasksReAssign(msg *message.Message) error {
	event := &TasksReAssign{}

	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		return errors.Wrap(err, "parse event")
	}

	return s.app.TasksReAssign(msg.Context(), *event)
}

func (s ServerInterfaceWrapper) UserCreated(msg *message.Message) error {
	event := &UserCreated{}

	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		return errors.Wrap(err, "parse event")
	}

	return s.app.UserCreated(msg.Context(), *event)
}

func (s ServerInterfaceWrapper) UserUpdated(msg *message.Message) error {
	event := &UserUpdated{}

	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		return errors.Wrap(err, "parse event")
	}

	return s.app.UserUpdated(msg.Context(), *event)
}