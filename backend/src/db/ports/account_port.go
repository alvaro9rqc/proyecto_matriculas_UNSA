package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type AccountRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (db.Account, error)
}
