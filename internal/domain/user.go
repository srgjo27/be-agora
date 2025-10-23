package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	AvatarURL    *string   `db:"avatar_url"`
	Role         string    `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
}
