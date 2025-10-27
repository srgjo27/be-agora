package domain

import (
	"time"

	"github.com/google/uuid"
)

type ThreadVote struct {
	UserID    uuid.UUID `db:"user_id"`
	ThreadID  uuid.UUID `db:"thread_id"`
	VoteType  int       `db:"vote_type"`
	CreatedAt time.Time `db:"created_at"`
}
