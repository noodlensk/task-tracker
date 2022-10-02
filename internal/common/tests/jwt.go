package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noodlensk/task-tracker/internal/common/auth"
)

func FakeAdminJWT(t *testing.T, userID string) string {
	t.Helper()

	token, err := auth.GetToken("secret", auth.User{
		UUID:  userID,
		Email: "admin@admin.com",
		Role:  "admin",
		Name:  "admin",
	})

	require.NoError(t, err)

	return token
}
