package usecase

import (
	"context"
	"errors"
	"serv_shop_haircompany/internal/modules/auth/v1/domain"
	dashboarduserdomen "serv_shop_haircompany/internal/modules/dashboard_user/v1/domain"
	"serv_shop_haircompany/internal/shared/infrastructure/security"
	"time"
)

type DashboardLoginUseCase interface {
	Execute(ctx context.Context, model *domain.DashboardLogin) (*security.TokenPair, error)
}

type dashboardLoginUseCase struct {
	repo              domain.SessionRepository
	dashboardUserRepo dashboarduserdomen.Repository
	tokenSvc          security.TokenService
}

func NewDashboardLoginUseCase(repo domain.SessionRepository, dashboardUserRepo dashboarduserdomen.Repository, tokenSvc security.TokenService) DashboardLoginUseCase {
	return &dashboardLoginUseCase{
		repo:              repo,
		dashboardUserRepo: dashboardUserRepo,
		tokenSvc:          tokenSvc,
	}
}

func (u *dashboardLoginUseCase) Execute(ctx context.Context, model *domain.DashboardLogin) (*security.TokenPair, error) {
	userModel, err := u.dashboardUserRepo.FindByEmail(ctx, model.Email.String())
	if err != nil {
		return nil, err
	}

	if !userModel.Password.Check(model.Password) {
		return nil, errors.New("invalid password")
	}

	tokens, jti, err := u.tokenSvc.GenerateTokenPair(userModel.ID, userModel.Role.String())
	if err != nil {
		return nil, err
	}

	sess := &domain.RefreshSession{
		JTI:       jti,
		UserID:    userModel.ID,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		Revoked:   false,
	}

	if err := u.repo.Save(ctx, sess, 30*24*time.Hour); err != nil {
		return nil, err
	}

	return tokens, nil
}
