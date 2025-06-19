package db

import (
	"github.com/enrollment/gen/db"
)

var StudentRepository *db.Queries = db.New(InstanceDB())
