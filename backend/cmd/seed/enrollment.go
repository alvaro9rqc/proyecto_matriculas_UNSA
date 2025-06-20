package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/enrollment/gen/db"
	"github.com/enrollment/src/db/ports"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"
	"golang.org/x/net/context"
)

const (
	VIRTUAL_MODALITY    string = "virtual"
	PRESENTIAL_MODALITY string = "presential"
	HYBRID_MODALITY     string = "hybrid"
)

var MODALITIES [3]string = [3]string{VIRTUAL_MODALITY, PRESENTIAL_MODALITY, HYBRID_MODALITY}

func createRandomMajor(faker faker.Faker) string {
	return strings.Join(faker.Lorem().Words(2), " ")
}

func createRandomCourse(faker faker.Faker) db.CreateCourseParams {
	return db.CreateCourseParams{
		Name:        strings.Join(faker.Lorem().Words(2), " "),
		Credits:     int16(rand.Intn(5) + 1),
		CicleNumber: int16(rand.Intn(5) + 1),
	}
}

func createRandomInstallation(faker faker.Faker) db.CreateInstalationParams {
	return db.CreateInstalationParams{
		Name: faker.Company().Name(),
		Description: pgtype.Text{
			String: faker.Lorem().Sentence(10),
			Valid:  true,
		},
	}
}

func createStudentGroup(priority int16) db.CreateStudentGroupParams {
	return db.CreateStudentGroupParams{
		Name:     fmt.Sprintf("Group %d", priority),
		Priority: priority,
		StartDay: pgtype.Date{
			Time:  time.Now().AddDate(0, 0, int(priority*7)),
			Valid: true,
		},
		EndDay: pgtype.Date{
			Time:  time.Now().AddDate(0, 0, int(priority*7)+6),
			Valid: true,
		},
	}
}

func seedEnrollmentRepository(ctx context.Context, studentGroupRepo ports.StudentGroupRepositoryInterface, installationRepo ports.InstallationRepositoryInterface, courseRepo ports.CourseRepositoryInterface, majorRepo ports.MajorRepositoryInterface, modalityRepo ports.ModalityRepositoryInterface) {
	faker := faker.New()

	log.Println("Seeding student groups...")

	for i := range 5 {
		studentGroup := createStudentGroup(int16(i + 1))
		err := studentGroupRepo.CreateStudentGroup(ctx, studentGroup)
		if err != nil {
			log.Fatalf("Failed to create student group: %v", err)
		}
	}

	log.Println("Seeding installations...")
	for range 40 {
		installation := createRandomInstallation(faker)
		err := installationRepo.CreateInstalation(ctx, installation)
		if err != nil {
			log.Fatalf("Failed to create installation: %v", err)
		}
	}

	log.Println("Seeding courses...")
	for range 100 {
		course := createRandomCourse(faker)
		err := courseRepo.CreateCourse(ctx, course)
		if err != nil {
			log.Fatalf("Failed to create course: %v", err)
		}
	}

	log.Println("Seeding majors...")
	for range 20 {
		majorName := createRandomMajor(faker)
		err := majorRepo.CreateMajor(ctx, majorName)
		if err != nil {
			log.Fatalf("Failed to create major: %v", err)
		}
	}

	log.Println("Seeding modalities...")
	for _, modality := range MODALITIES {
		err := modalityRepo.CreateModality(ctx, modality)
		if err != nil {
			log.Fatalf("Failed to create modality %s: %v", modality, err)
		}
	}

}
