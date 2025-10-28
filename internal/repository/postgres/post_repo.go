package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type postgresPostRepo struct {
	db *sqlx.DB
}

func NewPostgresPostRepo(db *sqlx.DB) usecase.PostRepository {
	return &postgresPostRepo{db: db}
}

func (r *postgresPostRepo) CountByThreadID(ctx context.Context, threadID uuid.UUID) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM posts WHERE thread_id = $1`
	err := r.db.GetContext(ctx, &count, query, threadID)

	return count, err
}

func (r *postgresPostRepo) Create(ctx context.Context, post *domain.Post) error {
	query := `INSERT INTO posts (id, content, user_id, thread_id, parent_post_id, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, post.ID, post.Content, post.UserID, post.ThreadID, post.ParentPostID, post.CreatedAt)

	return err
}

func (r *postgresPostRepo) GetByThreadID(ctx context.Context, threadID uuid.UUID, params usecase.PaginationParams) ([]*domain.Post, error) {
	var posts []*domain.Post
	query := `SELECT * 
	FROM posts 
	WHERE 
	thread_id = $1 
	ORDER BY created_at ASC 
	LIMIT $2 OFFSET $3`
	err := r.db.SelectContext(ctx, &posts, query, threadID, params.Limit, params.Offset)

	return posts, err
}

func (r *postgresPostRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	var post domain.Post
	query := `SELECT * FROM posts WHERE id = $1`
	err := r.db.GetContext(ctx, &post, query, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	return &post, err
}

func (r *postgresPostRepo) UpdateVoteCount(ctx context.Context, tx *sqlx.Tx, postID uuid.UUID, delta int) error {
	query := `UPDATE posts SET vote_count = vote_count + $1 WHERE id = $2`

	if tx != nil {
		_, err := tx.ExecContext(ctx, query, delta, postID)

		return err
	}

	_, err := r.db.ExecContext(ctx, query, delta, postID)

	return err
}
