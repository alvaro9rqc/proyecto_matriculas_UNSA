package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/enrollment/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

const hourLayout = "15:04:00"

var HOURS [16][2]string = [16][2]string{
	{"07:00:00", "07:50:00"},
	{"07:50:50", "08:40:00"},
	{"08:50:00", "09:40:00"},
	{"09:40:00", "10:30:00"},
	{"10:40:00", "11:30:00"},
	{"11:30:00", "12:20:00"},
	{"12:20:00", "13:10:00"},
	{"13:10:00", "14:00:00"},
	{"14:00:00", "14:50:00"},
	{"14:50:00", "15:40:00"},
	{"15:50:00", "16:40:00"},
	{"16:40:00", "17:30:00"},
	{"17:40:00", "18:30:00"},
	{"18:30:00", "19:20:00"},
	{"19:20:00", "20:10:00"},
	{"20:10:00", "21:00:00"},
}

var SECTIONS [5]string = [5]string{"A", "B", "C", "D", "E"}

func createRandomSection(course db.Course, section string) db.CreateSectionParams {
	return db.CreateSectionParams{
		CourseID: course.ID,
		Name:     section,
	}
}

func createRandomSlot(section db.Section) db.CreateSlotParams {
	return db.CreateSlotParams{
		SectionID:   section.ID,
		TotalPlaces: int32(30 + rand.Intn(5)),
		TakenPlaces: int32(5 + rand.Intn(20)),
	}
}

func createRandomEvent(section db.Section, installations []db.Installation, modalities []db.Modality, randomStartHour int) db.CreateEventParams {
	startTime := HOURS[randomStartHour][0]
	endTime := HOURS[randomStartHour][1]

	startTimeParsed, err := time.Parse(hourLayout, startTime)
	if err != nil {
		log.Fatalf("Failed to parse start time %s: %v", startTime, err)
	}
	endTimeParsed, err := time.Parse(hourLayout, endTime)
	if err != nil {
		log.Fatalf("Failed to parse end time %s: %v", endTime, err)
	}

	installation := installations[rand.Intn(len(installations))]
	modality := modalities[rand.Intn(len(modalities))]

	return db.CreateEventParams{
		StartDate:      pgtype.Timestamp{Time: startTimeParsed, Valid: true},
		EndDate:        pgtype.Timestamp{Time: endTimeParsed, Valid: true},
		SectionID:      section.ID,
		InstallationID: installation.ID,
		ModalityID:     modality.ID,
	}
}

func createRandomStudentMajorRelation(student db.FullListStudentsRow, major db.Major) db.CreateStudentMajorParams {
	return db.CreateStudentMajorParams{
		StudentID: student.ID,
		MajorID:   major.ID,
	}
}

func seedEnrollmentProcessTables(
	ctx context.Context,
	courseRepo ports.CourseRepositoryInterface,
	sectionRepo ports.SectionRepositoryInterface,
	slotsRepo ports.SlotsRepositoryInterface,
	installationRepo ports.InstallationRepositoryInterface,
	eventRepo ports.EventRepositoryInterface,
	modalityRepo ports.ModalityRepositoryInterface,
	studentRepo ports.StudentRepositoryInterface,
	majorRepo ports.MajorRepositoryInterface,
	studentMajorRepo ports.StudentMajorRepositoryInterface,
) {
	// faker := faker.New()

	log.Println("Seeding courses...")
	courses, err := courseRepo.ListCourses(ctx)
	if err != nil {
		log.Fatalf("Failed to list courses: %v", err)
	}
	for _, course := range courses {
		for _, section := range SECTIONS {
			sectionData := createRandomSection(course, section)
			err := sectionRepo.CreateSection(ctx, sectionData)
			if err != nil {
				log.Fatalf("Failed to create section %s for course %d: %v", section, course.ID, err)
			}
		}
	}

	log.Println("Seeding slots...")
	sections, err := sectionRepo.ListSections(ctx)
	if err != nil {
		log.Fatalf("Failed to list sections: %v", err)
	}
	for _, section := range sections {
		slotData := createRandomSlot(section)
		err := slotsRepo.CreateSlot(ctx, slotData)
		if err != nil {
			log.Fatalf("Failed to create slot for section %d: %v", section.ID, err)
		}
	}

	log.Println("Seeding events...")

	installationsAsync := <-utils.Async(func() ([]db.Installation, error) {
		return installationRepo.ListInstallations(ctx)
	})
	if installationsAsync.Err != nil {
		log.Fatalf("Failed to list installations: %v", installationsAsync.Err)
	}

	modalitiesAsync := <-utils.Async(func() ([]db.Modality, error) {
		return modalityRepo.ListModalities(ctx)
	})
	if modalitiesAsync.Err != nil {
		log.Fatalf("Failed to list modalities: %v", modalitiesAsync.Err)
	}

	installations := installationsAsync.Value
	modalities := modalitiesAsync.Value

	for _, section := range sections {
		numHours := (rand.Intn(3) + 1) // 1 to 3 blocks of hours, one block = 2 hours

		for range numHours {
			randomStartHour := rand.Intn(len(HOURS) - 1)

			event1 := createRandomEvent(section, installations, modalities, randomStartHour)
			event2 := createRandomEvent(section, installations, modalities, randomStartHour+1)

			err := eventRepo.CreateEvent(ctx, event1)
			if err != nil {
				log.Fatalf("Failed to create event for section %d: %v", section.ID, err)
			}

			err = eventRepo.CreateEvent(ctx, event2)
			if err != nil {
				log.Fatalf("Failed to create event for section %d: %v", section.ID, err)
			}
		}
	}

	log.Println("Seeding Student Major Relationships...")
	studentsAsync := <-utils.Async(func() ([]db.FullListStudentsRow, error) {
		return studentRepo.FullListStudents(ctx)
	})
	if studentsAsync.Err != nil {
		log.Fatalf("Failed to list students: %v", studentsAsync.Err)
	}

	majorsAsync := <-utils.Async(func() ([]db.Major, error) {
		return majorRepo.ListMajors(ctx)
	})
	if majorsAsync.Err != nil {
		log.Fatalf("Failed to list majors: %v", majorsAsync.Err)
	}

	students := studentsAsync.Value
	majors := majorsAsync.Value

	for _, student := range students {
		numOfMajorsByStudent := rand.Intn(2) + 1

		for range numOfMajorsByStudent {
			major := majors[rand.Intn(len(majors))]
			studentMajor := createRandomStudentMajorRelation(student, major)

			err := studentMajorRepo.CreateStudentMajor(ctx, studentMajor)
			if err != nil {
				log.Fatalf("Failed to create student-major relation for student %d and major %d: %v", student.ID, major.ID, err)
			}
		}
	}

}
