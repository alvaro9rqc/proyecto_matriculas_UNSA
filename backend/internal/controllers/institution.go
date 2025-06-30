package controllers

import (
	"context"
	"fmt"

	institution "github.com/enrollment/gen/institution"
	"github.com/enrollment/internal/ports"
	"github.com/enrollment/internal/utils"
)

type institutionsrvc struct {
	OauthRepo       ports.OauthRepositoryInterface
	InstitutionRepo ports.InstitutionRepositoryInterface
}

func NewInstitution(oauthRepo ports.OauthRepositoryInterface, institutionRepo ports.InstitutionRepositoryInterface) institution.Service {
	return &institutionsrvc{
		OauthRepo:       oauthRepo,
		InstitutionRepo: institutionRepo,
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
