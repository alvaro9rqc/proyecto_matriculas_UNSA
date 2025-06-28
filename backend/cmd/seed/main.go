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
		oauthRepo          = repositories.NewOauthRepository(conn)
		studentGroupRepo   = repositories.NewStudentGroupRepository(conn)
		installationRepo   = repositories.NewInstallationRepository(conn)
		courseRepo         = repositories.NewCourseRepository(conn)
		processRepo        = repositories.NewProcessRepository(conn)
		modalityRepo       = repositories.NewModalityRepository(conn)
		studentRepo        = repositories.NewStudentRepository(conn)
		speakerRepo        = repositories.NewSpeakerRepository(conn)
		studentProcessRepo = repositories.NewStudentProcessRepository(conn)
		sectionSpeakerRepo = repositories.NewSectionSpeakerRepository(conn)
		sectionRepo        = repositories.NewSectionRepository(conn)
		slotsRepo          = repositories.NewSlotsRepository(conn)
		eventRepo          = repositories.NewEventRepository(conn)
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
		processRepo,
		modalityRepo,
		oauthRepo,
		studentRepo,
		speakerRepo,
	)
	seedEnrollmentProcessTables(
		cfg.Ctx,
		courseRepo,
		sectionRepo,
		slotsRepo,
		installationRepo,
		eventRepo,
		modalityRepo,
		studentRepo,
		processRepo,
		studentProcessRepo,
		speakerRepo,
		sectionSpeakerRepo,
	)
}
