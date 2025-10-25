package domain

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	ID         uuid.UUID  `db:"id"`
	Title      string     `db:"title"`
	Slug       string     `db:"slug"`
	Content    string     `db:"content"`
	UserID     uuid.UUID  `db:"user_id"`
	CategoryID uuid.UUID  `db:"category_id"`
	IsPinned   bool       `db:"is_pinned"`
	IsLocked   bool       `db:"is_locked"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}
