package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type postgresCategoryRepo struct {
	db *sqlx.DB
}

func (r *postgresCategoryRepo) GetByIDs(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]*domain.Category, error) {
	categories := []*domain.Category{}
	query, args, err := sqlx.In(`SELECT id, name, slug, description, created_at FROM categories WHERE id IN (?)`, ids)

	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	err = r.db.SelectContext(ctx, &categories, query, args...)
	if err != nil {
		return nil, err
	}

	catMap := make(map[uuid.UUID]*domain.Category)
	for _, cat := range categories {
		catMap[cat.ID] = cat
	}

	return catMap, nil
}

func NewPostgresCategoryRepo(db *sqlx.DB) usecase.CategoryRepository {
	return &postgresCategoryRepo{db: db}
}

func (r *postgresCategoryRepo) Create(ctx context.Context, category *domain.Category) error {
	query := `INSERT INTO categories (id, name, slug, description, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, category.ID, category.Name, category.Slug, category.Description, category.CreatedAt)

	return err
}

func (r *postgresCategoryRepo) GetBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	var category domain.Category

	query := `SELECT id, name, slug, description, created_at FROM categories WHERE slug = $1`

	err := r.db.GetContext(ctx, &category, query, slug)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	return &category, err
}

func (r *postgresCategoryRepo) GetAll(ctx context.Context) ([]*domain.Category, error) {
	var categories []*domain.Category

	query := `SELECT * FROM categories ORDER BY created_at ASC`

	err := r.db.SelectContext(ctx, &categories, query)

	return categories, err
}

func (r *postgresCategoryRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	var category domain.Category

	query := `SELECT * FROM categories WHERE id = $1`

	err := r.db.GetContext(ctx, &category, query, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}

	return &category, err
}
