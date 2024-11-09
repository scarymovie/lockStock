package repository

import (
	"context"
	"database/sql"
	"lockStock/internal/domain/user"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) SaveUser(ctx context.Context, user *user.User) (string, error) {
	query := `INSERT INTO users (uid, created_at) VALUES ($1, $2) RETURNING id`
	err := u.db.QueryRowContext(ctx, query, user.UID, user.CreatedAt).Scan(&user.UID)
	return user.UID, err
}
