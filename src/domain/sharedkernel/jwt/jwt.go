package jwt

import "context"

type JWT interface {
	GenerateAccessToken(ctx context.Context, userId string) (*string, error)
	GenerateRefreshToken(ctx context.Context, userId string) (*string, error)
}
