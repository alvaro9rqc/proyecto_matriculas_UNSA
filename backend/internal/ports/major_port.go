package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type MajorRepositoryInterface interface {
	CreateMajor(ctx context.Context, name string) error
	ListMajors(ctx context.Context) ([]db.Major, error)
	DeleteMajor(ctx context.Context, id int32) error
	UpdateMajor(ctx context.Context, arg db.UpdateMajorParams) error
}
