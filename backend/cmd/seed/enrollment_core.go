package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"
	"golang.org/x/net/context"
)

const (
	VIRTUAL_MODALITY    string = "virtual"
	PRESENTIAL_MODALITY string = "presential"
	HYBRID_MODALITY     string = "hybrid"
)

const (
	NUM_STUDENTS int = NUM_ACCOUNTS - 10
)

var MODALITIES [3]string = [3]string{VIRTUAL_MODALITY, PRESENTIAL_MODALITY, HYBRID_MODALITY}

func createRandomStudent(faker faker.Faker, account db.Account, studentGroup db.StudentGroup) (db.CreateStudentParams, error) {
	return db.CreateStudentParams{
		AccountID:      account.ID,
		Code:           faker.Person().SSN(),
		StudentGroupID: studentGroup.ID,
	}, nil
}

func createRandomMajor(faker faker.Faker) string {
	return strings.Join(faker.Lorem().Words(2), " ")
}

func createRandomCourse(faker faker.Faker, major db.Major) db.CreateCourseParams {
	return db.CreateCourseParams{
		Name:        strings.Join(faker.Lorem().Words(2), " "),
		Credits:     int16(rand.Intn(5) + 1),
		CycleNumber: int16(rand.Intn(5) + 1),
		MajorID:     major.ID,
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

func seedEnrollmentCoreTables(
	ctx context.Context,
	studentGroupRepo ports.StudentGroupRepositoryInterface,
	installationRepo ports.InstallationRepositoryInterface,
	courseRepo ports.CourseRepositoryInterface,
	majorRepo ports.MajorRepositoryInterface,
	modalityRepo ports.ModalityRepositoryInterface,
	oauthRepo ports.OauthRepositoryInterface,
	studentRepo ports.StudentRepositoryInterface,
	speakerRepo ports.SpeakerRepositoryInterface,
) {
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

	log.Println("Seeding majors...")
	for range 20 {
		majorName := createRandomMajor(faker)
		err := majorRepo.CreateMajor(ctx, majorName)
		if err != nil {
			log.Fatalf("Failed to create major: %v", err)
		}
	}

	log.Println("Seeding courses...")
	majors, err := majorRepo.ListMajors(ctx)
	if err != nil {
		log.Fatalf("Failed to list majors: %v", err)
	}
	for range 100 {
		major := majors[rand.Intn(len(majors))]
		course := createRandomCourse(faker, major)
		err := courseRepo.CreateCourse(ctx, course)
		if err != nil {
			log.Fatalf("Failed to create course: %v", err)
		}
	}

	log.Println("Seeding modalities...")
	for _, modality := range MODALITIES {
		err := modalityRepo.CreateModality(ctx, modality)
		if err != nil {
			log.Fatalf("Failed to create modality %s: %v", modality, err)
		}
	}

	log.Print("Seeding students...")
	accounts, err := oauthRepo.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  int32(NUM_STUDENTS),
		Offset: 0,
	})
	if err != nil {
		log.Fatalf("Failed to list accounts: %v", err)
	}
	studentGroups, err := studentGroupRepo.ListStudentGroups(ctx)
	if err != nil {
		log.Fatalf("Failed to list student groups: %v", err)
	}
	for i, account := range accounts {
		studentGroup := studentGroups[i%len(studentGroups)]
		student, err := createRandomStudent(faker, account, studentGroup)
		if err != nil {
			log.Fatalf("Failed to create student: %v", err)
		}
		err = studentRepo.CreateStudent(ctx, student)
		if err != nil {
			log.Fatalf("Failed to create student for account %d: %v", account.ID, err)
		}
	}

	log.Println("Seeding skeapers...")
	accounts, err = oauthRepo.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  int32(NUM_ACCOUNTS),
		Offset: int32(NUM_STUDENTS),
	})
	if err != nil {
		log.Fatalf("Failed to list accounts: %v", err)
	}
	for _, account := range accounts {
		err = speakerRepo.CreateSpeaker(ctx, account.ID)
		if err != nil {
			log.Fatalf("Failed to create skeaper for account %d: %v", account.ID, err)
		}
	}
}
