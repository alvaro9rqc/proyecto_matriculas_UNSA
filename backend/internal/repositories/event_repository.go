package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type EventRepository struct {
	*db.Queries
}

func NewEventRepository(conn *pgx.Conn) ports.EventRepositoryInterface {
	return &EventRepository{
		Queries: db.New(conn),
	}
}
