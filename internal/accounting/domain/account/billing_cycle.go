package account

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type BillingCycle struct {
	uid      string
	userUIID string
	from     time.Time
	to       time.Time
	isClosed bool
}

func (c *BillingCycle) UID() string               { return c.uid }
func (c *BillingCycle) UserUID() string           { return c.userUIID }
func (c *BillingCycle) From() time.Time           { return c.from }
func (c *BillingCycle) To() time.Time             { return c.to }
func (c *BillingCycle) IsClosed() bool            { return c.isClosed }
func (c *BillingCycle) IsActual(t time.Time) bool { return c.from.After(t) && c.to.After(t) }

func (c *BillingCycle) Close() error {
	if c.isClosed {
		return errors.New("already closed")
	}

	c.isClosed = true

	return nil
}

func NewBillingCycleForDate(t time.Time) BillingCycle {
	return BillingCycle{
		uid:      uuid.New().String(),
		userUIID: "",
		from:     time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, t.Nanosecond(), t.Location()),
		to:       time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, t.Nanosecond(), t.Location()),
		isClosed: false,
	}
}
