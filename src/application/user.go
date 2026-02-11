package application

import (
	"context"
	"time"
	errorDomain "tlab/src/domain/error"
	"tlab/src/domain/sharedkernel/jwt"
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/domain/sharedkernel/unitofwork"
	"tlab/src/domain/user"
	"tlab/src/infrastructure/http/dto"

	"github.com/google/uuid"
)

type User struct {
	uow      unitofwork.UnitOfWork
	jwt      jwt.JWT
	logger   logger.Logger
	userRepo user.UserRepository
}

func NewUser(
	uow unitofwork.UnitOfWork,
	jwt jwt.JWT,
	logger logger.Logger,
	userRepo user.UserRepository,

) *User {
	return &User{
		uow:      uow,
		jwt:      jwt,
		logger:   logger,
		userRepo: userRepo,
	}
}

func (a *User) Login(ctx context.Context, payload dto.LoginInput) (*user.JWTToken, error) {
	creds := user.LoginInput{
		Email:    payload.Email,
		Password: payload.Password,
	}

	user, err := creds.GetUser(ctx, a.userRepo)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errorDomain.ErrUserNotFound
	}

	if err = user.CheckPassword(payload.Password); err != nil {
		return nil, err
	}

	token, err := user.GenerateJWTToken(ctx, a.jwt)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *User) Register(ctx context.Context, payload dto.RegisterInput) error {
	userId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	user := &user.User{
		ID:        userId,
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	salt, err := user.GenerateSalt()
	if err != nil {
		return err
	}

	user.Password = string(salt)

	duplicate, err := user.IsEmailDuplicate(ctx, a.userRepo)
	if err != nil || duplicate {
		return err
	}

	return user.NewUser(ctx, a.userRepo)
}
