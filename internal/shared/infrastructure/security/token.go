package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AccessClaims struct {
	UserID uint   `json:"uid"`
	Role   string `json:"role,omitempty"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID uint   `json:"uid"`
	JTI    string `json:"jti"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenService interface {
	GenerateTokenPair(userID uint, role string) (*TokenPair, string, error)
	ParseAccessToken(tokenStr string) (*AccessClaims, error)
	ParseRefreshToken(tokenStr string) (*RefreshClaims, error)
}

type jwtService struct {
	secretKey  []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewTokenService(secret string) TokenService {
	return &jwtService{
		secretKey:  []byte(secret),
		accessTTL:  time.Hour,           // 1 час
		refreshTTL: 30 * 24 * time.Hour, // 30 дней
	}
}

func (s *jwtService) GenerateTokenPair(userID uint, role string) (*TokenPair, string, error) {
	now := time.Now()
	jti := uuid.NewString()

	accessClaims := AccessClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "serv_shop_haircompany",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTTL)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString(s.secretKey)
	if err != nil {
		return nil, "", err
	}

	refreshClaims := RefreshClaims{
		UserID: userID,
		JTI:    jti,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "serv_shop_haircompany",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTTL)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString(s.secretKey)
	if err != nil {
		return nil, "", err
	}

	return &TokenPair{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
	}, jti, nil
}

func (s *jwtService) ParseAccessToken(tokenStr string) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&AccessClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return s.secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	return claims, nil
}

func (s *jwtService) ParseRefreshToken(tokenStr string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&RefreshClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return s.secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}
