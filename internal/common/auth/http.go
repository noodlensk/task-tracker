package auth

import (
	"context"
	"net/http"
	"strings"

	commonerrors "github.com/noodlensk/task-tracker/internal/common/errors"
	"github.com/noodlensk/task-tracker/internal/common/server/httperr"
)

type JWTHttpMiddleware struct {
	Secret  string
	AuthURL string
}

func (a JWTHttpMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == a.AuthURL { // do not require token for auth endpoint
			next.ServeHTTP(w, r)

			return
		}

		ctx := r.Context()

		bearerToken := a.tokenFromHeader(r)
		if bearerToken == "" {
			httperr.Unauthorized("empty-bearer-token", nil, w, r)
			return
		}

		user, err := ParseToken(a.Secret, bearerToken)
		if err != nil {
			httperr.Unauthorized("unable-to-verify-jwt", err, w, r)
			return
		}

		ctx = context.WithValue(ctx, userContextKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (a JWTHttpMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

var (
	// if we expect that the user of the function may be interested with concrete error,
	// it's a good idea to provide variable with this error
	NoUserInContextError = commonerrors.NewAuthorizationError("no user in context", "no-user-found")
)

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}
