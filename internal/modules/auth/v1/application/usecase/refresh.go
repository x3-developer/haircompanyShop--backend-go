package usecase

import (
	"context"
	"errors"
	"serv_shop_haircompany/internal/modules/auth/v1/domain"
	"serv_shop_haircompany/internal/shared/infrastructure/security"
	"time"
)

type RefreshUseCase interface {
	Execute(ctx context.Context, refreshToken string) (*security.TokenPair, error)
}

type refreshUseCase struct {
	repo     domain.SessionRepository
	tokenSvc security.TokenService
}

func NewRefreshUseCase(repo domain.SessionRepository, tokenSvc security.TokenService) RefreshUseCase {
	return &refreshUseCase{
		repo:     repo,
		tokenSvc: tokenSvc,
	}
}

func (u refreshUseCase) Execute(ctx context.Context, refreshToken string) (*security.TokenPair, error) {
	claims, err := u.tokenSvc.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	sess, err := u.repo.Get(ctx, claims.UserID, claims.JTI)
	if err != nil || sess == nil || sess.Revoked {
		return nil, errors.New("refresh session not found or revoked")
	}

	err = u.repo.Delete(ctx, claims.UserID, claims.JTI)
	if err != nil {
		return nil, err
	}

	tokens, newJTI, err := u.tokenSvc.GenerateTokenPair(claims.UserID, claims.Role)
	if err != nil {
		return nil, err
	}

	newSess := &domain.RefreshSession{
		JTI:       newJTI,
		UserID:    claims.UserID,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		Revoked:   false,
	}

	if err := u.repo.Save(ctx, newSess, 30*24*time.Hour); err != nil {
		return nil, err
	}

	return tokens, nil
}
