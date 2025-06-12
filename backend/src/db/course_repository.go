package db

import (
	"github.com/enrollment/gen/db"
)

var CourseRepository *db.Queries = db.New(InstanceDB())
