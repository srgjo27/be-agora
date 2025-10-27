package http

import (
	"time"

	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/domain"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	AvatarURL *string   `json:"avatar_url"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type CategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewCategoryResponse(cat *domain.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:          cat.ID,
		Name:        cat.Name,
		Slug:        cat.Slug,
		Description: cat.Description,
		CreatedAt:   cat.CreatedAt,
	}
}

func NewCategoryListResponse(cats []*domain.Category) []*CategoryResponse {
	list := make([]*CategoryResponse, len(cats))
	for i, cat := range cats {
		list[i] = NewCategoryResponse(cat)
	}

	return list
}

type ThreadSummaryResponse struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	UserID     uuid.UUID `json:"user_id"`
	CategoryID uuid.UUID `json:"category_id"`
	IsPinned   bool      `json:"is_pinned"`
	IsLocked   bool      `json:"is_locked"`
	VoteCount  int       `json:"vote_count"`
	CreatedAt  time.Time `json:"created_at"`
}

type ThreadDetailResponse struct {
	ID         uuid.UUID  `json:"id"`
	Title      string     `json:"title"`
	Slug       string     `json:"slug"`
	Content    string     `json:"content"`
	UserID     uuid.UUID  `json:"user_id"`
	CategoryID uuid.UUID  `json:"category_id"`
	IsPinned   bool       `json:"is_pinned"`
	IsLocked   bool       `json:"is_locked"`
	VoteCount  int        `json:"vote_count"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

func NewThreadDetailResponse(t *domain.Thread) *ThreadDetailResponse {
	return &ThreadDetailResponse{
		ID:         t.ID,
		Title:      t.Title,
		Slug:       t.Slug,
		Content:    t.Content,
		UserID:     t.UserID,
		CategoryID: t.CategoryID,
		IsPinned:   t.IsPinned,
		IsLocked:   t.IsLocked,
		VoteCount:  t.VoteCount,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
	}
}

func NewThreadSummaryResponse(t *domain.Thread) *ThreadSummaryResponse {
	return &ThreadSummaryResponse{
		ID:         t.ID,
		Title:      t.Title,
		Slug:       t.Slug,
		UserID:     t.UserID,
		CategoryID: t.CategoryID,
		IsPinned:   t.IsPinned,
		IsLocked:   t.IsLocked,
		VoteCount:  t.VoteCount,
		CreatedAt:  t.CreatedAt,
	}
}

func NewThreadListResponse(threads []*domain.Thread) []*ThreadSummaryResponse {
	list := make([]*ThreadSummaryResponse, len(threads))
	for i, t := range threads {
		list[i] = NewThreadSummaryResponse(t)
	}

	return list
}

type PostResponse struct {
	ID           uuid.UUID  `json:"id"`
	Content      string     `json:"content"`
	UserID       uuid.UUID  `json:"user_id"`
	ThreadID     uuid.UUID  `json:"thread_id"`
	ParentPostID *uuid.UUID `json:"parent_post_id,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

func NewPostResponse(p *domain.Post) *PostResponse {
	return &PostResponse{
		ID:           p.ID,
		Content:      p.Content,
		UserID:       p.UserID,
		ThreadID:     p.ThreadID,
		ParentPostID: p.ParentPostID,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func NewPostListResponse(posts []*domain.Post) []*PostResponse {
	list := make([]*PostResponse, len(posts))
	for i, p := range posts {
		list[i] = NewPostResponse(p)
	}

	return list
}
