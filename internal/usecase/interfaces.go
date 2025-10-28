package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srgjo27/agora/internal/domain"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByIDs(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]*domain.User, error)
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
	GetByIDs(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]*domain.Category, error)
}

type CategoryUsecase interface {
	Create(ctx context.Context, name string, description *string) (*domain.Category, error)
	GetAll(ctx context.Context) ([]*domain.Category, error)
}

type UpdateThreadParams struct {
	Title   *string
	Content *string
}

type ThreadRepository interface {
	Create(ctx context.Context, thread *domain.Thread) error
	GetAll(ctx context.Context, params PaginationParams) ([]*domain.Thread, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Thread, error)
	UpdateVoteCount(ctx context.Context, tx *sqlx.Tx, threadID uuid.UUID, delta int) error
	CountAll(ctx context.Context) (int, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, thread *domain.Thread) error
}

type ThreadUsecase interface {
	Create(ctx context.Context, title, content string, userID, categoryID uuid.UUID) (*domain.Thread, *domain.User, *domain.Category, error)
	GetAll(ctx context.Context, params PaginationParams) ([]*domain.Thread, map[uuid.UUID]*domain.User, map[uuid.UUID]*domain.Category, int, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Thread, *domain.User, *domain.Category, error)
	Delete(ctx context.Context, threadID, userID uuid.UUID, role string) error
	Update(ctx context.Context, threadID, userID uuid.UUID, role string, params UpdateThreadParams) (*domain.Thread, *domain.User, *domain.Category, error)
}

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	GetByThreadID(ctx context.Context, threadID uuid.UUID, params PaginationParams) ([]*domain.Post, error)
	CountByThreadID(ctx context.Context, threadID uuid.UUID) (int, error)
}

type PostUsecase interface {
	Create(ctx context.Context, content string, userID, threadID uuid.UUID, parentPostID *uuid.UUID) (*domain.Post, error)
	GetByThreadID(ctx context.Context, threadID uuid.UUID, params PaginationParams) ([]*domain.Post, int, error)
}

type VoteRepository interface {
	GetThreadVote(ctx context.Context, userID, threadID uuid.UUID) (*domain.ThreadVote, error)
	UpsertThreadVote(ctx context.Context, tx *sqlx.Tx, vote *domain.ThreadVote) error
	DeleteThreadVote(ctx context.Context, tx *sqlx.Tx, userID, threadID uuid.UUID) error
}

type VoteUsecase interface {
	VoteOnThread(ctx context.Context, userID, threadID uuid.UUID, voteType int) error
}

type PaginationParams struct {
	Limit  int
	Offset int
}
