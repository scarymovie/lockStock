package contract

import (
	"context"
	"lockStock/internal/domain/user"
)

type UserCreator interface {
	CreateUser(ctx context.Context, newUser *user.User) (string, error)
}
