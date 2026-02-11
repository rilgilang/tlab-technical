package user

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, payload User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id string) (*User, error)
}
