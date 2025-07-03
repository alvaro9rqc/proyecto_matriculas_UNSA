package controllers

import (
	"context"
	"fmt"

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
}

func NewInstitution(oauthRepo ports.OauthRepositoryInterface, institutionRepo ports.InstitutionRepositoryInterface, processRepo ports.ProcessRepositoryInterface, studentRepo ports.StudentRepositoryInterface, courseRepo ports.CourseRepositoryInterface) institution.Service {
	return &institutionsrvc{
		OauthRepo:       oauthRepo,
		InstitutionRepo: institutionRepo,
		ProcessRepo:     processRepo,
		StudentRepo:     studentRepo,
		CourseRepo:      courseRepo,
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

func (s *institutionsrvc) ListProccesByInstitution(ctx context.Context, payload *institution.ListProccesByInstitutionPayload) (res []*institution.Process, err error) {
	token := utils.GetTokenFromContext(ctx)

	session, err := s.OauthRepo.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, institution.MakeNotAuthorized(fmt.Errorf("failed to get session by token: %w", err))
	}

	processes, err := s.ProcessRepo.ListProcessByInstitutionId(ctx, db.ListProcessByInstitutionIdParams{
		ID:            session.AccountID,
		InstitutionID: payload.InstitutionID,
	})
	if err != nil {
		return nil, institution.MakeInternalServerError(fmt.Errorf("failed to list processes: %w", err))
	}

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

	return processesByInstitution, nil
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
