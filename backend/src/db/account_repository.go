package db

import (
	"github.com/enrollment/gen/db"
)

type AccountRepository struct {
	*db.Queries
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		Queries: db.New(InstanceDB()),
	}
}
