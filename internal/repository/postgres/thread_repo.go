package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type postgresThreadRepo struct {
	db *sqlx.DB
}

func NewPostgresThreadRepo(db *sqlx.DB) usecase.ThreadRepository {
	return &postgresThreadRepo{db: db}
}

func (r *postgresThreadRepo) Create(ctx context.Context, thread *domain.Thread) error {
	query := `INSERT INTO threads (id, title, slug, content, user_id, category_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, thread.ID, thread.Title, thread.Slug, thread.Content, thread.UserID, thread.CategoryID, thread.CreatedAt)

	return err
}

func (r *postgresThreadRepo) GetAll(ctx context.Context, params usecase.PaginationParams) ([]*domain.Thread, error) {
	var threads []*domain.Thread

	query := `SELECT * FROM threads ORDER BY is_pinned DESC, created_at DESC LIMIT $1 OFFSET $2`
	err := r.db.SelectContext(ctx, &threads, query, params.Limit, params.Offset)

	return threads, err
}

func (r *postgresThreadRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Thread, error) {
	var thread domain.Thread

	query := `SELECT * FROM threads WHERE id = $1`

	err := r.db.GetContext(ctx, &thread, query, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	return &thread, err
}

func (r *postgresThreadRepo) UpdateVoteCount(ctx context.Context, tx *sqlx.Tx, threadID uuid.UUID, delta int) error {
	query := `UPDATE threads SET vote_count = vote_count + $1 WHERE id = $2`
	if tx != nil {
		_, err := tx.ExecContext(ctx, query, delta, threadID)
		return err
	}

	_, err := r.db.ExecContext(ctx, query, delta, threadID)

	return err
}

func (r *postgresThreadRepo) CountAll(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM threads`
	err := r.db.GetContext(ctx, &count, query)
	return count, err
}

func (r *postgresThreadRepo) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM threads WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *postgresThreadRepo) Update(ctx context.Context, thread *domain.Thread) error {
	query := `UPDATE threads SET title = $1, slug = $2, content = $3, updated_at = $4 WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query, thread.Title, thread.Slug, thread.Content, thread.UpdatedAt, thread.ID)

	return err
}
