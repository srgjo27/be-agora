package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID           uuid.UUID  `db:"id"`
	Content      string     `db:"content"`
	UserID       uuid.UUID  `db:"user_id"`
	ThreadID     uuid.UUID  `db:"thread_id"`
	ParentPostID *uuid.UUID `db:"parent_post_id"`
	VoteCount    int        `db:"vote_count"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}
