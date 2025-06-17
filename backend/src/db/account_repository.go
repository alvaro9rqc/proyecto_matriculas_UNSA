package db

import (
	"github.com/enrollment/gen/db"
)

var AccountRepository *db.Queries = db.New(InstanceDB())
