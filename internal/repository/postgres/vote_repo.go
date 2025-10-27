package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type postgresVoteRepo struct {
	db *sqlx.DB
}

func (r *postgresVoteRepo) DeleteThreadVote(ctx context.Context, tx *sqlx.Tx, userID uuid.UUID, threadID uuid.UUID) error {
	query := `DELETE FROM thread_votes WHERE user_id = $1 AND thread_id = $2`
	if tx != nil {
		_, err := tx.ExecContext(ctx, query, userID, threadID)
		return err
	}

	_, err := r.db.ExecContext(ctx, query, userID, threadID)

	return err
}

func (r *postgresVoteRepo) GetThreadVote(ctx context.Context, userID uuid.UUID, threadID uuid.UUID) (*domain.ThreadVote, error) {
	var vote domain.ThreadVote

	query := `SELECT user_id, thread_id, vote_type, created_at FROM thread_votes WHERE user_id = $1 AND thread_id = $2`

	err := r.db.GetContext(ctx, &vote, query, userID, threadID)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	return &vote, err
}

func (r *postgresVoteRepo) UpsertThreadVote(ctx context.Context, tx *sqlx.Tx, vote *domain.ThreadVote) error {
	query := `INSERT INTO thread_votes (user_id, thread_id, vote_type, created_at) VALUES ($1, $2, $3, $4) ON CONFLICT (user_id, thread_id) DO UPDATE SET vote_type = EXCLUDED.vote_type`

	vote.CreatedAt = time.Now()

	if tx != nil {
		_, err := tx.ExecContext(ctx, query, vote.UserID, vote.ThreadID, vote.VoteType, vote.CreatedAt)

		return err
	}

	_, err := r.db.ExecContext(ctx, query, vote.UserID, vote.ThreadID, vote.VoteType, vote.CreatedAt)
	return err
}

func NewPostgresVoteRepo(db *sqlx.DB) usecase.VoteRepository {
	return &postgresVoteRepo{db: db}
}
