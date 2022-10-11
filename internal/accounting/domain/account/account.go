package account

import (
	"fmt"
	"time"
)

type Account struct {
	userUID       string
	balance       float32
	billingCycles []BillingCycle
	transactions  []Transaction
}

func (a *Account) UserUID() string               { return a.userUID }
func (a *Account) Balance() float32              { return a.balance }
func (a *Account) Withdraw(amount float32)       { a.balance -= amount }
func (a *Account) TopUp(amount float32)          { a.balance += amount }
func (a *Account) BillingCycles() []BillingCycle { return a.billingCycles }
func (a *Account) Transactions() []Transaction   { return a.transactions }
func (a *Account) AddTransaction(t *Transaction) error {
	var bc BillingCycle

	for _, c := range a.billingCycles {
		if !c.isClosed && c.IsActual(time.Now()) {
			bc = c

			break
		}
	}

	if bc.UID() == "" { // TODO: check for null point ref
		bc = NewBillingCycleForDate(time.Now())
	}

	t.billingCycleUID = bc.UID()
	a.transactions = append(a.transactions, *t)

	a.balance += t.debit
	a.balance -= t.credit

	return nil
}

func (a *Account) CloseBillingCycle() error {
	for _, bc := range a.billingCycles {
		if bc.isClosed {
			continue
		}

		total := float32(0)

		for _, t := range a.transactions {
			if t.billingCycleUID != bc.UID() {
				continue
			}

			total += t.Debit()
			total += t.Credit()
		}

		if total > 0 {
			reason := fmt.Sprintf("Payout for %s - %s", bc.From(), bc.To())

			payOutTransaction, err := NewCreditTransaction(bc.UserUID(), total, reason)
			if err != nil {
				return err
			}

			payOutTransaction.billingCycleUID = bc.UID()

			if err := a.AddTransaction(payOutTransaction); err != nil {
				return err
			}

			a.Withdraw(total)
		}

		if err := bc.Close(); err != nil {
			return err
		}
	}

	return nil
}
