package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/config"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type jwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

type tokenService struct {
	cfg *config.Config
}

func NewTokenService(cfg *config.Config) usecase.TokenService {
	return &tokenService{cfg: cfg}
}

func (s *tokenService) GenerateRefreshToken(ctx context.Context, user *domain.User) (string, error) {
	expirationTime := time.Now().Add(s.cfg.RefreshTokenDurationHours * time.Hour)

	claims := &jwtClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "agora-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.cfg.JWTSecretKey))
}

func (s *tokenService) GenerateAccessToken(ctx context.Context, user *domain.User) (string, error) {
	expirationTime := time.Now().Add(s.cfg.AccessTokenDurationMinutes * time.Minute)

	claims := &jwtClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "agora-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.cfg.JWTSecretKey))
}
