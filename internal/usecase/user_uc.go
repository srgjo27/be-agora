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
	tokenSvc TokenService
}

func NewUserUsecase(ur UserRepository, ts TokenService) UserUsecase {
	return &userUsecase{userRepo: ur, tokenSvc: ts}
}

func (uc *userUsecase) Login(ctx context.Context, email string, password string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", domain.ErrInvalid
	}

	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if err == domain.ErrNotFound {
			return "", "", domain.ErrUnauthorized
		}

		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", domain.ErrUnauthorized
	}

	accessToken, err := uc.tokenSvc.GenerateAccessToken(ctx, user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := uc.tokenSvc.GenerateRefreshToken(ctx, user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (uc *userUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return uc.userRepo.GetByID(ctx, id)
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

func (uc *userUsecase) Refresh(ctx context.Context, refreshToken string) (string, error) {
	userID, _, err := uc.tokenSvc.ValidateToken(ctx, refreshToken)

	if err != nil {
		return "", domain.ErrUnauthorized
	}

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		if err == domain.ErrNotFound {
			return "", domain.ErrUnauthorized
		}

		return "", err
	}

	newAccessToken, err := uc.tokenSvc.GenerateAccessToken(ctx, user)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (uc *userUsecase) GetUsers(ctx context.Context) ([]*domain.User, error) {
	return uc.userRepo.GetUsers(ctx)
}
