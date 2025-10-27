package http

import "github.com/google/uuid"

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateCategoryRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
}

type CreateThreadRequest struct {
	Title      string    `json:"title" binding:"required,min=5"`
	Content    string    `json:"content" binding:"required,min=10"`
	CategoryID uuid.UUID `json:"category_id" binding:"required"`
}

type CreatePostRequest struct {
	Content      string     `json:"content" binding:"required"`
	ParentPostID *uuid.UUID `json:"parent_post_id"`
}
