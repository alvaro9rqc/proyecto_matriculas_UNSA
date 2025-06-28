package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type ProcessRepositoryInterface interface {
	CreateProcess(ctx context.Context, args db.CreateProcessParams) error
	ListAllProcess(ctx context.Context) ([]db.Process, error)
	DeleteProcess(ctx context.Context, id int32) error
}
