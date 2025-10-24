package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/srgjo27/agora/internal/domain"
)

type categoryUsecase struct {
	categoryRepo CategoryRepository
}

func NewCategoryUsecase(cr CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo: cr}
}

func (uc *categoryUsecase) Create(ctx context.Context, name string, description *string) (*domain.Category, error) {
	if name == "" {
		return nil, domain.ErrInvalid
	}

	categorySlug := slug.Make(name)

	existing, err := uc.categoryRepo.GetBySlug(ctx, categorySlug)
	if err != nil && err != domain.ErrNotFound {
		return nil, err
	}

	if existing != nil {
		return nil, domain.ErrConflict
	}

	category := &domain.Category{
		ID:          uuid.New(),
		Name:        name,
		Slug:        categorySlug,
		Description: description,
		CreatedAt:   time.Now(),
	}

	if err := uc.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (uc *categoryUsecase) GetAll(ctx context.Context) ([]*domain.Category, error) {
	return uc.categoryRepo.GetAll(ctx)
}
