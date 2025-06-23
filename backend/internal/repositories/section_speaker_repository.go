package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type SectionSpeakerRepository struct {
	*db.Queries
}

func NewSectionSpeakerRepository(conn *pgx.Conn) ports.SectionSpeakerRepositoryInterface {
	return &SectionSpeakerRepository{
		Queries: db.New(conn),
	}
}
