package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/domain"
)

type postUsecase struct {
	postRepo   PostRepository
	threadRepo ThreadRepository
}

func (uc *postUsecase) Create(ctx context.Context, content string, userID uuid.UUID, threadID uuid.UUID, parentPostID *uuid.UUID) (*domain.Post, error) {
	if content == "" {
		return nil, domain.ErrInvalid
	}

	thread, err := uc.threadRepo.GetByID(ctx, threadID)
	if err != nil {
		if err == domain.ErrNotFound {
			return nil, domain.ErrInvalid
		}

		return nil, err
	}

	if thread.IsLocked {
		return nil, domain.ErrThreadLocked
	}

	// (optional) validate parentPostID
	if parentPostID != nil {
		// You might want to implement a check here to see if the parent post exists
	}

	post := &domain.Post{
		ID:           uuid.New(),
		Content:      content,
		UserID:       userID,
		ThreadID:     threadID,
		ParentPostID: parentPostID,
		CreatedAt:    time.Now(),
	}

	if err := uc.postRepo.Create(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *postUsecase) GetByThreadID(ctx context.Context, threadID uuid.UUID) ([]*domain.Post, error) {
	_, err := uc.threadRepo.GetByID(ctx, threadID)
	if err != nil {
		return nil, err
	}

	return uc.postRepo.GetByThreadID(ctx, threadID)
}

func NewPostUsecase(pr PostRepository, tr ThreadRepository) PostUsecase {
	return &postUsecase{
		postRepo:   pr,
		threadRepo: tr,
	}
}
