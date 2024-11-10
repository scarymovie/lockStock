package contract

import (
	"context"
	"database/sql"
	"lockStock/internal/domain/user"
)

type UserCreator interface {
	CreateUser(ctx context.Context, tx *sql.Tx, newUser *user.User) (string, error)
}
