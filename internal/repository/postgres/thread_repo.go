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

func (r *postgresThreadRepo) GetAll(ctx context.Context) ([]*domain.Thread, error) {
	var threads []*domain.Thread

	query := `SELECT * FROM threads ORDER BY is_pinned DESC, created_at DESC`
	err := r.db.SelectContext(ctx, &threads, query)

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
