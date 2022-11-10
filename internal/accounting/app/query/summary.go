package query

import (
	"context"

	"github.com/noodlensk/task-tracker/internal/accounting/domain/account"
)

type SummaryParams struct{}

type Summary struct{}

type SummaryHandler struct {
	accountRepo account.Repository
}

func (h *SummaryHandler) Handle(ctx context.Context, q UserAccountParams) (*Summary, error) {
	return nil, nil
}
