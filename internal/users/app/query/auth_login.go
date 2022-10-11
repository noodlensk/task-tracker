package query

import (
	"context"
	"errors"

	"github.com/noodlensk/task-tracker/internal/common/auth"
	"github.com/noodlensk/task-tracker/internal/users/domain/user"
)

type AuthLogin struct {
	Email    string
	Password string
}

type AuthResult struct {
	Token string
}

type AuthLoginHandler struct {
	repo   user.Repository
	secret string
}

func NewAuthLoginHandler(repo user.Repository, secret string) AuthLoginHandler {
	return AuthLoginHandler{repo: repo, secret: secret}
}

func (h AuthLoginHandler) Handle(ctx context.Context, q AuthLogin) (*AuthResult, error) {
	u, err := h.repo.GetUserByEmail(ctx, q.Email)
	if err != nil {
		return nil, err
	}

	if !u.PasswordMatch(q.Password) {
		return nil, errors.New("password not match") // TODO: replace with proper type
	}

	token, err := auth.GetToken(h.secret, auth.User{
		UUID:  u.UID(),
		Email: u.Email(),
		Role:  u.Role().String(),
		Name:  u.Name(),
	})
	if err != nil {
		return nil, err
	}

	return &AuthResult{Token: token}, nil
}
