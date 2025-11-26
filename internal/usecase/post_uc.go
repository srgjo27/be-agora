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
	userRepo   UserRepository
}

func NewPostUsecase(pr PostRepository, tr ThreadRepository, ur UserRepository) PostUsecase {
	return &postUsecase{
		postRepo:   pr,
		threadRepo: tr,
		userRepo:   ur,
	}
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

func (uc *postUsecase) GetByThreadID(ctx context.Context, threadID uuid.UUID, params PaginationParams) ([]*domain.Post, map[uuid.UUID]*domain.User, int, error) {
	_, err := uc.threadRepo.GetByID(ctx, threadID)
	if err != nil {
		return nil, nil, 0, err
	}

	total, err := uc.postRepo.CountByThreadID(ctx, threadID)
	if err != nil {
		return nil, nil, 0, err
	}

	posts, err := uc.postRepo.GetByThreadID(ctx, threadID, params)
	if err != nil {
		return nil, nil, 0, err
	}

	if len(posts) == 0 {
		return []*domain.Post{}, map[uuid.UUID]*domain.User{}, total, nil
	}

	userIDs := make([]uuid.UUID, 0)
	for _, p := range posts {
		userIDs = append(userIDs, p.UserID)
	}

	userMap, err := uc.userRepo.GetByIDs(ctx, userIDs)
	if err != nil {
		return nil, nil, 0, err
	}

	return posts, userMap, total, nil
}
