package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/domain"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type UserUsecase interface {
	Register(ctx context.Context, username, email, password string) (*domain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Login(ctx context.Context, email, password string) (accessToken string, refreshToken string, err error)
	Refresh(ctx context.Context, refreshToken string) (newAccessToken string, err error)
}

type TokenService interface {
	GenerateAccessToken(ctx context.Context, user *domain.User) (string, error)
	GenerateRefreshToken(ctx context.Context, user *domain.User) (string, error)
	ValidateToken(ctx context.Context, tokenString string) (uuid.UUID, string, error)
}

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	GetBySlug(ctx context.Context, slug string) (*domain.Category, error)
	GetAll(ctx context.Context) ([]*domain.Category, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Category, error)
}

type CategoryUsecase interface {
	Create(ctx context.Context, name string, description *string) (*domain.Category, error)
	GetAll(ctx context.Context) ([]*domain.Category, error)
}

type ThreadRepository interface {
	Create(ctx context.Context, thread *domain.Thread) error
	GetAll(ctx context.Context) ([]*domain.Thread, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Thread, error)
}

type ThreadUsecase interface {
	Create(ctx context.Context, title, content string, userID, categoryID uuid.UUID) (*domain.Thread, error)
	GetAll(ctx context.Context) ([]*domain.Thread, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Thread, error)
}
