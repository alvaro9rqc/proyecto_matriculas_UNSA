package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type AccountSessionRepository interface {
	CreateAccountSession(ctx context.Context, arg db.CreateAccountSessionParams) (db.AccountSession, error)
}
