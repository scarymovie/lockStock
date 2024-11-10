package service

import (
	"context"
	"database/sql"
	"lockStock/internal/domain/user"
)

// UserRepository работает с данными пользователей.
type UserRepository struct {
	tx *sql.Tx
}

// NewUserRepository создаёт новый репозиторий с транзакцией.
func NewUserRepository(tx *sql.Tx) *UserRepository {
	return &UserRepository{tx: tx}
}

// SaveUser сохраняет пользователя в базе данных с использованием транзакции.
func (r *UserRepository) SaveUser(ctx context.Context, newUser *user.User) (string, error) {
	query := `INSERT INTO users (uid, created_at) VALUES ($1, $2) RETURNING id`
	var id string
	err := r.tx.QueryRowContext(ctx, query, newUser.UID, newUser.CreatedAt).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
