package tests

import (
	"context"
	"fmt"
	"net/http"
)

func authorizationBearer(token string) func(context.Context, *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		return nil
	}
}
