package task

import (
	"errors"
)

type Task struct {
	uid               string
	title             string
	assignedToUserUID string
	status            string
	priceAssigned     float32
	priceFinished     float32
}

func (t *Task) UID() string               { return t.uid }
func (t *Task) Title() string             { return t.title }
func (t *Task) AssignedToUserUID() string { return t.assignedToUserUID }
func (t *Task) Status() string            { return t.status }
func (t *Task) PriceAssigned() float32    { return t.priceAssigned }
func (t *Task) PriceFinished() float32    { return t.priceFinished }
func (t *Task) SetPriceAssigned(price float32) error {
	if price < 0 {
		return errors.New("price should not be < 0")
	}

	t.priceAssigned = price

	return nil
}

func (t *Task) SetPriceFinished(price float32) error {
	if price < 0 {
		return errors.New("price should not be < 0")
	}

	t.priceFinished = price

	return nil
}

func (t *Task) SetTitle(title string)   { t.title = title }
func (t *Task) SetStatus(status string) { t.status = status }
func (t *Task) AssignToUser(uid string) { t.assignedToUserUID = uid }

func NewTaskFromCUD(uid, title, assignedToUserUID, status string) Task {
	return Task{
		uid:               uid,
		title:             title,
		assignedToUserUID: assignedToUserUID,
		status:            status,
	}
}
