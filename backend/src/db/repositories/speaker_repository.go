package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/src/db/ports"
	"github.com/jackc/pgx/v5"
)

type SpeakerRepository struct {
	*db.Queries
}

func NewSpeakerRepository(conn *pgx.Conn) ports.SpeakerRepositoryInterface {
	return &SpeakerRepository{
		Queries: db.New(conn),
	}
}
