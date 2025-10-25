package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/srgjo27/agora/internal/domain"
)

type threadUsecase struct {
	threadRepo   ThreadRepository
	categoryRepo CategoryRepository
}

func NewThreadUsecase(tr ThreadRepository, cr CategoryRepository) ThreadUsecase {
	return &threadUsecase{
		threadRepo:   tr,
		categoryRepo: cr,
	}
}

func (uc *threadUsecase) Create(ctx context.Context, title string, content string, userID uuid.UUID, categoryID uuid.UUID) (*domain.Thread, error) {
	if title == "" || content == "" {
		return nil, domain.ErrInvalid
	}

	_, err := uc.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, domain.ErrInvalid
		}

		return nil, err
	}

	threadSlug := slug.Make(title)

	thread := &domain.Thread{
		ID:         uuid.New(),
		Title:      title,
		Slug:       threadSlug,
		Content:    content,
		UserID:     userID,
		CategoryID: categoryID,
		CreatedAt:  time.Now(),
	}

	if err := uc.threadRepo.Create(ctx, thread); err != nil {
		return nil, err
	}

	return thread, nil
}

func (uc *threadUsecase) GetAll(ctx context.Context) ([]*domain.Thread, error) {
	return uc.threadRepo.GetAll(ctx)
}

func (uc *threadUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Thread, error) {
	return uc.threadRepo.GetByID(ctx, id)
}
