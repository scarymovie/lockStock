package contract

import (
	"context"
	"lockStock/internal/domain/user"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *user.User) (string, error)
}
