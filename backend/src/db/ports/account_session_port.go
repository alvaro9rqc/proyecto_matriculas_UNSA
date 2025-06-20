package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type AccountSessionRepositoryInterface interface {
	CreateAccountSession(ctx context.Context, arg db.CreateAccountSessionParams) (db.AccountSession, error)
}
