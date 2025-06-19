package db

import (
	"github.com/enrollment/gen/db"
)

type AccountSessionRepository struct {
	*db.Queries
}

func NewAccountSessionRepository() *AccountSessionRepository {
	return &AccountSessionRepository{
		Queries: db.New(InstanceDB()),
	}
}
