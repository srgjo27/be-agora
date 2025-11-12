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
	ID        uuid.UUID             `json:"id"`
	Title     string                `json:"title"`
	Slug      string                `json:"slug"`
	Author    *AuthorResponse       `json:"author"`
	Category  *CategoryInfoResponse `json:"category"`
	IsPinned  bool                  `json:"is_pinned"`
	IsLocked  bool                  `json:"is_locked"`
	VoteCount int                   `json:"vote_count"`
	CreatedAt time.Time             `json:"created_at"`
}

type ThreadDetailResponse struct {
	ID        uuid.UUID             `json:"id"`
	Title     string                `json:"title"`
	Slug      string                `json:"slug"`
	Content   string                `json:"content"`
	Author    *AuthorResponse       `json:"author"`
	Category  *CategoryInfoResponse `json:"category"`
	IsPinned  bool                  `json:"is_pinned"`
	IsLocked  bool                  `json:"is_locked"`
	VoteCount int                   `json:"vote_count"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt *time.Time            `json:"updated_at,omitempty"`
}

func NewThreadDetailResponse(t *domain.Thread, author *domain.User, cat *domain.Category) *ThreadDetailResponse {
	return &ThreadDetailResponse{
		ID:        t.ID,
		Title:     t.Title,
		Slug:      t.Slug,
		Content:   t.Content,
		Author:    NewAuthorResponse(author),
		Category:  NewCategoryInfoResponse(cat),
		IsPinned:  t.IsPinned,
		IsLocked:  t.IsLocked,
		VoteCount: t.VoteCount,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func NewThreadSummaryResponse(t *domain.Thread, author *domain.User, cat *domain.Category) *ThreadSummaryResponse {
	return &ThreadSummaryResponse{
		ID:        t.ID,
		Title:     t.Title,
		Slug:      t.Slug,
		Author:    NewAuthorResponse(author),
		Category:  NewCategoryInfoResponse(cat),
		IsPinned:  t.IsPinned,
		IsLocked:  t.IsLocked,
		VoteCount: t.VoteCount,
		CreatedAt: t.CreatedAt,
	}
}

type PostResponse struct {
	ID           uuid.UUID       `json:"id"`
	Content      string          `json:"content"`
	Author       *AuthorResponse `json:"author"`
	ThreadID     uuid.UUID       `json:"thread_id"`
	ParentPostID *uuid.UUID      `json:"parent_post_id,omitempty"`
	VoteCount    int             `json:"vote_count"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    *time.Time      `json:"updated_at,omitempty"`
}

func NewPostResponse(p *domain.Post, author *domain.User) *PostResponse {
	return &PostResponse{
		ID:           p.ID,
		Content:      p.Content,
		Author:       NewAuthorResponse(author),
		ThreadID:     p.ThreadID,
		ParentPostID: p.ParentPostID,
		VoteCount:    p.VoteCount,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

type PaginationMeta struct {
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
	Limit       int `json:"limit"`
}

type AuthorResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	AvatarURL *string   `json:"avatar_url"`
}

type CategoryInfoResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}

func NewAuthorResponse(user *domain.User) *AuthorResponse {
	if user == nil {
		return nil
	}

	return &AuthorResponse{
		ID:        user.ID,
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
	}
}

func NewCategoryInfoResponse(cat *domain.Category) *CategoryInfoResponse {
	if cat == nil {
		return nil
	}

	return &CategoryInfoResponse{
		ID:   cat.ID,
		Name: cat.Name,
		Slug: cat.Slug,
	}
}
