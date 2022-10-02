package tests

import (
	"github.com/noodlensk/task-tracker/internal/common/auth"
	"github.com/stretchr/testify/require"
	"testing"
)

func FakeAdminJWT(t *testing.T, userID string) string {
	token, err := auth.GetToken("secret", auth.User{
		UUID:  userID,
		Email: "admin@admin.com",
		Role:  "admin",
		Name:  "admin",
	})

	require.NoError(t, err)

	return token
}
