package usecase

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/srgjo27/agora/internal/domain"
)

type threadUsecase struct {
	threadRepo   ThreadRepository
	categoryRepo CategoryRepository
	userRepo     UserRepository
}

func NewThreadUsecase(tr ThreadRepository, cr CategoryRepository, ur UserRepository) ThreadUsecase {
	return &threadUsecase{
		threadRepo:   tr,
		categoryRepo: cr,
		userRepo:     ur,
	}
}

func (uc *threadUsecase) Create(ctx context.Context, title string, content string, userID uuid.UUID, categoryID uuid.UUID) (*domain.Thread, *domain.User, *domain.Category, error) {
	if title == "" || content == "" {
		return nil, nil, nil, domain.ErrInvalid
	}

	category, err := uc.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, nil, nil, domain.ErrInvalid
		}

		return nil, nil, nil, err
	}

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, nil, nil, domain.ErrInvalid
		}

		return nil, nil, nil, err
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
		VoteCount:  0,
	}

	if err := uc.threadRepo.Create(ctx, thread); err != nil {
		return nil, nil, nil, err
	}

	return thread, user, category, nil
}

func (uc *threadUsecase) GetAll(ctx context.Context, params PaginationParams) ([]*domain.Thread, map[uuid.UUID]*domain.User, map[uuid.UUID]*domain.Category, int, error) {
	threads, err := uc.threadRepo.GetAll(ctx, params)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	total, err := uc.threadRepo.CountAll(ctx)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	if len(threads) == 0 {
		return []*domain.Thread{}, nil, nil, total, nil
	}

	userIDs := make([]uuid.UUID, 0)
	catIDs := make([]uuid.UUID, 0)
	for _, t := range threads {
		userIDs = append(userIDs, t.UserID)
		catIDs = append(catIDs, t.CategoryID)
	}

	userMap, err := uc.userRepo.GetByIDs(ctx, userIDs)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	catMap, err := uc.categoryRepo.GetByIDs(ctx, catIDs)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	return threads, userMap, catMap, total, nil
}

func (uc *threadUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Thread, *domain.User, *domain.Category, error) {
	thread, err := uc.threadRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	user, err := uc.userRepo.GetByID(ctx, thread.UserID)
	if err != nil {
		log.Printf("[ERROR]: User not found for thread %s: %v", id, err)
	}

	cat, err := uc.categoryRepo.GetByID(ctx, thread.CategoryID)
	if err != nil {
		log.Printf("[ERROR]: Category not found for thread %s: %v", id, err)
	}

	return thread, user, cat, nil
}
