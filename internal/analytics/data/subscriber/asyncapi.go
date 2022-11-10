package subscriber

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/analytics/domain/account"
	"github.com/noodlensk/task-tracker/internal/analytics/domain/user"
)

type TransactionRepository interface {
	StoreTransaction(ctx context.Context, t account.Transaction) error
}

type UsersRepository interface {
	StoreUser(ctx context.Context, u user.User) error
}

type Server struct {
	transactionsRepo TransactionRepository
	usersRepo        UsersRepository
}

func (s Server) TransactionCreated(ctx context.Context, e TransactionCreated) error {
	if err := s.transactionsRepo.StoreTransaction(ctx, account.Transaction{
		UID:     e.Uid,
		UserUID: e.UserUid,
		Amount:  e.Amount,
		Reason:  e.Reason,
	}); err != nil {
		return err
	}

	return nil
}

func (s Server) UserCreated(ctx context.Context, e UserCreated) error {
	return s.usersRepo.StoreUser(ctx, user.User{
		UID:   e.Id,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	})
}

func (s Server) UserUpdated(ctx context.Context, e UserUpdated) error {
	return s.usersRepo.StoreUser(ctx, user.User{
		UID:   e.Id,
		Name:  e.Name,
		Email: e.Email,
		Role:  e.Role,
	})
}

func NewAsyncServer(transactionRepo TransactionRepository, usersRepo UsersRepository) *Server {
	return &Server{transactionsRepo: transactionRepo, usersRepo: usersRepo}
}
