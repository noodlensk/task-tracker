package account

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	uid             string
	userUID         string
	billingCycleUID string
	debit           float32
	credit          float32
	reason          string
	createdAt       time.Time
}

func (t *Transaction) UID() string             { return t.uid }
func (t *Transaction) UserUID() string         { return t.userUID }
func (t *Transaction) BillingCycleUID() string { return t.billingCycleUID }
func (t *Transaction) Debit() float32          { return t.debit }
func (t *Transaction) Credit() float32         { return t.credit }
func (t *Transaction) Reason() string          { return t.reason }
func (t *Transaction) CreatedAt() time.Time    { return t.createdAt }

func NewDebitTransaction(userUID string, debit float32, reason string) (*Transaction, error) {
	if userUID == "" {
		return nil, errors.New("userUID can't be empty")
	}

	if debit < 0 {
		return nil, errors.New("debit can't be < 0")
	}

	if reason == "" {
		return nil, errors.New("reason can't be empty")
	}

	return &Transaction{
		uid:       uuid.New().String(),
		userUID:   userUID,
		debit:     debit,
		reason:    reason,
		createdAt: time.Now(),
	}, nil
}

func NewCreditTransaction(userUID string, credit float32, reason string) (*Transaction, error) {
	if userUID == "" {
		return nil, errors.New("userUID can't be empty")
	}

	if credit < 0 {
		return nil, errors.New("credit can't be < 0")
	}

	if reason == "" {
		return nil, errors.New("reason can't be empty")
	}

	return &Transaction{
		uid:       uuid.New().String(),
		userUID:   userUID,
		credit:    credit,
		reason:    reason,
		createdAt: time.Now(),
	}, nil
}
