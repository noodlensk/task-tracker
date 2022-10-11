package task

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/noodlensk/task-tracker/internal/tasks/domain/user"
)

type Task struct {
	uid         string
	title       string
	description string
	assignedTo  user.User
	status      Status
	createdBy   user.User
	createdAt   time.Time
	updatedAt   time.Time
}

type Status string

const (
	StatusOpen = Status("open")
	StatusDone = Status("done")
)

func (t *Task) Assign(u user.User) {
	t.assignedTo = u
	t.updatedAt = time.Now()
}

func (t *Task) Complete() {
	t.status = StatusDone
	t.updatedAt = time.Now()
}

func (t *Task) Title() string         { return t.title }
func (t *Task) Description() string   { return t.description }
func (t *Task) Status() Status        { return t.status }
func (t *Task) AssignedTo() user.User { return t.assignedTo }
func (t *Task) CreatedBy() user.User  { return t.createdBy }
func (t *Task) CreatedAt() time.Time  { return t.createdAt }
func (t *Task) UpdatedAt() time.Time  { return t.updatedAt }
func (t *Task) UID() string           { return t.uid }

func NewTask(title, desc string, createdBy user.User) (*Task, error) {
	if title == "" {
		return nil, errors.New("empty title")
	}

	if desc == "" {
		return nil, errors.New("empty description")
	}

	t := &Task{
		uid:         uuid.New().String(),
		title:       title,
		description: desc,
		status:      StatusOpen,
		createdBy:   createdBy,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	return t, nil
}
