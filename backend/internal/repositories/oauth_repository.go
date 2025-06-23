package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type OauthRepository struct {
	*db.Queries
}

func NewOauthRepository(conn *pgx.Conn) ports.OauthRepositoryInterface {
	return &OauthRepository{
		Queries: db.New(conn),
	}
}
