package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/enrollment/gen/db"
	institution "github.com/enrollment/gen/institution"
	"github.com/enrollment/internal/ports"
	"github.com/enrollment/internal/utils"
)

type institutionsrvc struct {
	OauthRepo       ports.OauthRepositoryInterface
	InstitutionRepo ports.InstitutionRepositoryInterface
	ProcessRepo     ports.ProcessRepositoryInterface
	StudentRepo     ports.StudentRepositoryInterface
	CourseRepo      ports.CourseRepositoryInterface
	SectionRepo     ports.SectionRepositoryInterface
}

func NewInstitution(oauthRepo ports.OauthRepositoryInterface, institutionRepo ports.InstitutionRepositoryInterface, processRepo ports.ProcessRepositoryInterface, studentRepo ports.StudentRepositoryInterface, courseRepo ports.CourseRepositoryInterface, sectionRepo ports.SectionRepositoryInterface) institution.Service {
	return &institutionsrvc{
		OauthRepo:       oauthRepo,
		InstitutionRepo: institutionRepo,
		ProcessRepo:     processRepo,
		StudentRepo:     studentRepo,
		CourseRepo:      courseRepo,
		SectionRepo:     sectionRepo,
	}
}

func (s *institutionsrvc) ListInstitutions(ctx context.Context) (res []*institution.Institution, err error) {
	token := utils.GetTokenFromContext(ctx)

	session, err := s.OauthRepo.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, institution.MakeNotAuthorized(fmt.Errorf("failed to get session by token: %w", err))
	}

	institutions, err := s.InstitutionRepo.ListInstitutionsByAccountID(ctx, session.AccountID)
	if err != nil {
		return nil, institution.MakeInternalServerError(fmt.Errorf("failed to list institutions: %w", err))
	}

	institutionsByAccount := make([]*institution.Institution, 0, len(institutions))

	for _, inst := range institutions {
		institutionsByAccount = append(institutionsByAccount, &institution.Institution{
			ID:      inst.ID,
			Name:    inst.Name,
			LogoURL: &inst.LogoUrl.String,
		})
	}

	return institutionsByAccount, nil
}

func (s *institutionsrvc) ListProccesByInstitution(ctx context.Context, payload *institution.ListProccesByInstitutionPayload) (res *institution.ListProccesByInstitutionResult, err error) {
	token := utils.GetTokenFromContext(ctx)

	session, err := s.OauthRepo.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, institution.MakeNotAuthorized(fmt.Errorf("failed to get session by token: %w", err))
	}

	instChan := utils.Async(func() (db.Institution, error) {
		return s.InstitutionRepo.GetInstitutionByID(ctx, payload.InstitutionID)
	})
	processesChan := utils.Async(func() ([]db.Process, error) {
		return s.ProcessRepo.ListProcessByInstitutionId(ctx, db.ListProcessByInstitutionIdParams{
			ID:            session.AccountID,
			InstitutionID: payload.InstitutionID,
		})
	})

	instAsync := <-instChan
	processesAsync := <-processesChan

	if instAsync.Err != nil {
		return nil, institution.MakeInternalServerError(fmt.Errorf("failed to get institution by ID: %w", instAsync.Err))
	}
	inst := instAsync.Value
	if processesAsync.Err != nil {
		return nil, institution.MakeInternalServerError(fmt.Errorf("failed to list processes: %w", processesAsync.Err))
	}
	processes := processesAsync.Value

	processesByInstitution := make([]*institution.Process, 0, len(processes))
	for _, proc := range processes {
		startDay := proc.StartDay.Time.Unix()
		endDay := proc.EndDay.Time.Unix()

		processesByInstitution = append(processesByInstitution, &institution.Process{
			ID:            &proc.ID,
			Name:          &proc.Name,
			StartDay:      &startDay,
			EndDay:        &endDay,
			InstitutionID: &proc.InstitutionID,
		})
	}

	return &institution.ListProccesByInstitutionResult{
		Processes: processesByInstitution,
		ID:        inst.ID,
		Name:      inst.Name,
		LogoURL:   &inst.LogoUrl.String,
	}, nil
}

// List all courses available for a student in a specific process
func (s *institutionsrvc) ListAllCoursesAvailableByStudentInProcess(ctx context.Context, p *institution.ListAllCoursesAvailableByStudentInProcessPayload) (res []*institution.Course, err error) {
	token := utils.GetTokenFromContext(ctx)
	_, err = s.OauthRepo.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, institution.MakeNotAuthorized(fmt.Errorf("failed to get session by token: %w", err))
	}

	studentId, err := s.StudentRepo.GetStudentIdByToken(ctx, token)

	raw_courses, err := s.CourseRepo.ListAllCoursesAvailableByStudentInProcess(ctx, db.ListAllCoursesAvailableByStudentInProcessParams{
		StudentID: studentId,
		ProcessID: p.ProcessID,
	})

	if err != nil {
		return nil, institution.MakeInternalServerError(fmt.Errorf("failed to list courses: %w", err))
	}

	coursesAvailable := make([]*institution.Course, 0, len(raw_courses))
	for _, course := range raw_courses {
		coursesAvailable = append(coursesAvailable, &institution.Course{
			ID:          int(course.CourseID),
			Name:        course.CourseName,
			Credits:     int(course.Credits),
			CicleNumber: int(course.CycleNumber),
		})
	}

	return coursesAvailable, nil
}

// Expand a course to get detailed information about their events and sections
func (s *institutionsrvc) ExpandCourse(ctx context.Context, p *institution.ExpandCoursePayload) (res []*institution.SectionWithEvents, err error) {
	token := utils.GetTokenFromContext(ctx)
	_, err = s.OauthRepo.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, institution.MakeNotAuthorized(fmt.Errorf("failed to get session by token: %w", err))
	}

	map_sec_idx := map[int32]int{}

	rows, err := s.SectionRepo.ListDetailedSectionByCourseId(ctx, p.CourseID)

	for _, row := range rows {
		idx, exits := map_sec_idx[row.SectionID]
		if !exits {
			idx = len(res)
			map_sec_idx[row.SectionID] = idx
			res = append(res, &institution.SectionWithEvents{
				ID:          int(row.SectionID),
				SectionName: row.SectionName,
				TakenPlaces: int(row.TakenPlaces),
				TotalPlaces: int(row.TotalPlaces),
				Events:      []*institution.DetailedEvent{},
			})
		}
		if row.EventID.Valid {
			res[idx].Events = append(res[idx].Events, &institution.DetailedEvent{
				ID:               int(row.EventID.Int32),
				StartDate:        row.StartDate.Time.Format(time.RFC3339),
				EndDate:          row.EndDate.Time.Format(time.RFC3339),
				SectionID:        int(row.SectionID),
				InstallationID:   int(row.InstallationID.Int32),
				InstallationName: row.InstallationName.String,
				ModalityID:       int(row.ModalityID.Int32),
				ModalityName:     row.ModalityName.String,
			})
		}
	}
	return
}
