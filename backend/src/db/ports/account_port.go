package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type AccountRepositoryInterface interface {
	GetAccountByEmail(ctx context.Context, email string) (db.Account, error)
}
