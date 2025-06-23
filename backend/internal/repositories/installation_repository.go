package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type InstallationRepository struct {
	*db.Queries
}

func NewInstallationRepository(conn *pgx.Conn) ports.InstallationRepositoryInterface {
	return &InstallationRepository{
		Queries: db.New(conn),
	}
}
