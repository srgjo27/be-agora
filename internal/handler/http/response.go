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
