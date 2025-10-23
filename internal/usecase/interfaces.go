package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/domain"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
}

type UserUsecase interface {
	Register(ctx context.Context, username, email, password string) (*domain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Login(ctx context.Context, email, password string) (accessToken string, refreshToken string, err error)
}

type TokenService interface {
	GenerateAccessToken(ctx context.Context, user *domain.User) (string, error)
	GenerateRefreshToken(ctx context.Context, user *domain.User) (string, error)
}
