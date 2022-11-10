package adapters

import (
	"context"
	"sync"

	"github.com/noodlensk/task-tracker/internal/analytics/domain/account"
)

type TransactionInMemoryRepository struct {
	transactions map[string]*account.Transaction

	lock sync.RWMutex
}

func NewTransactionInMemoryRepository() *TransactionInMemoryRepository {
	return &TransactionInMemoryRepository{transactions: map[string]*account.Transaction{}}
}

func (r *TransactionInMemoryRepository) StoreTransaction(ctx context.Context, t account.Transaction) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.transactions[t.UID] = &t

	return nil
}
