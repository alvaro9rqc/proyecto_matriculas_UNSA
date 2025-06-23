package main

import (
	"github.com/enrollment/config"
	"github.com/enrollment/internal/db"
	"github.com/enrollment/internal/repositories"
)

func main() {
	cfg, err := config.NewSeedConfig()
	if err != nil {
		panic(err)
	}

	conn, err := db.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}

	var (
		oauthRepo        = repositories.NewOauthRepository(conn)
		studentGroupRepo = repositories.NewStudentGroupRepository(conn)
		installationRepo = repositories.NewInstallationRepository(conn)
		courseRepo       = repositories.NewCourseRepository(conn)
		majorRepo        = repositories.NewMajorRepository(conn)
		modalityRepo     = repositories.NewModalityRepository(conn)
		studentRepo      = repositories.NewStudentRepository(conn)
		speakerRepo      = repositories.NewSpeakerRepository(conn)
	)

	seedOauthTables(
		cfg.Ctx,
		oauthRepo,
	)
	seedEnrollmentCoreTables(
		cfg.Ctx,
		studentGroupRepo,
		installationRepo,
		courseRepo,
		majorRepo,
		modalityRepo,
		oauthRepo,
		studentRepo,
		speakerRepo,
	)
	// seedEnrollmentProcessTables(
	// )
}
