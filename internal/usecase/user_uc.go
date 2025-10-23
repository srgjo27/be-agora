package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo UserRepository
}

func (uc *userUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return nil, nil
}

func (uc *userUsecase) Register(ctx context.Context, username string, email string, password string) (*domain.User, error) {
	if email == "" || password == "" || username == "" {
		return nil, domain.ErrInvalid
	}

	existingUser, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil && err != domain.ErrNotFound {
		return nil, err
	}

	if existingUser != nil {
		return nil, domain.ErrConflict
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:           uuid.New(),
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         "member",
		CreatedAt:    time.Now(),
	}

	err = uc.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserUsecase(ur UserRepository) UserUsecase {
	return &userUsecase{userRepo: ur}
}
