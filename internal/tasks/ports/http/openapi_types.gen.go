// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package http

import (
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for TaskStatus.
const (
	DONE TaskStatus = "DONE"
	NEW  TaskStatus = "NEW"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// Task defines model for Task.
type Task struct {
	AssignedTo  *string     `json:"assigned_to,omitempty"`
	CreatedAt   *time.Time  `json:"created_at,omitempty"`
	CreatedBy   *string     `json:"created_by,omitempty"`
	Description string      `json:"description"`
	ModifiedAt  *time.Time  `json:"modified_at,omitempty"`
	Status      *TaskStatus `json:"status,omitempty"`
	Title       string      `json:"title"`
	Uid         *string     `json:"uid,omitempty"`
}

// TaskStatus defines model for Task.Status.
type TaskStatus string

// TaskUpdate defines model for TaskUpdate.
type TaskUpdate struct {
	Uid []string `json:"uid"`
}

// Tasks defines model for Tasks.
type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// GetTasksParams defines parameters for GetTasks.
type GetTasksParams struct {
	// The number of items to skip before starting to collect the result set
	Offset int `form:"offset" json:"offset"`

	// The numbers of items to return
	Limit int `form:"limit" json:"limit"`
}

// CreateTaskJSONBody defines parameters for CreateTask.
type CreateTaskJSONBody = Task

// MarkTaskAsCompleteJSONBody defines parameters for MarkTaskAsComplete.
type MarkTaskAsCompleteJSONBody = TaskUpdate

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody = CreateTaskJSONBody

// MarkTaskAsCompleteJSONRequestBody defines body for MarkTaskAsComplete for application/json ContentType.
type MarkTaskAsCompleteJSONRequestBody = MarkTaskAsCompleteJSONBody
