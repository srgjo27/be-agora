package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type postgresPostRepo struct {
	db *sqlx.DB
}

func (r *postgresPostRepo) Create(ctx context.Context, post *domain.Post) error {
	query := `INSERT INTO posts (id, content, user_id, thread_id, parent_post_id, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, post.ID, post.Content, post.UserID, post.ThreadID, post.ParentPostID, post.CreatedAt)

	return err
}

func (r *postgresPostRepo) GetByThreadID(ctx context.Context, threadID uuid.UUID) ([]*domain.Post, error) {
	var posts []*domain.Post
	query := `SELECT * FROM posts WHERE thread_id = $1 ORDER BY created_at ASC`
	err := r.db.SelectContext(ctx, &posts, query, threadID)

	return posts, err
}

func NewPostgresPostRepo(db *sqlx.DB) usecase.PostRepository {
	return &postgresPostRepo{db: db}
}
