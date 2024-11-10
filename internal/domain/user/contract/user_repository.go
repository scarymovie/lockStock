package contract

import (
	"context"
	"database/sql"
	"lockStock/internal/domain/user"
)

type UserRepository interface {
	SaveUser(ctx context.Context, tx *sql.Tx, user *user.User) (string, error)
}
