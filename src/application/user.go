package application

import (
	"context"
	"time"
	errorDomain "tlab/src/domain/error"
	"tlab/src/domain/sharedkernel/jwt"
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/domain/sharedkernel/unitofwork"
	"tlab/src/domain/user"
	"tlab/src/domain/wallet"
	"tlab/src/infrastructure/http/dto"

	"github.com/google/uuid"
)

type User struct {
	uow        unitofwork.UnitOfWork
	jwt        jwt.JWT
	logger     logger.Logger
	userRepo   user.UserRepository
	walletRepo wallet.WalletRepository
}

func NewUser(
	uow unitofwork.UnitOfWork,
	jwt jwt.JWT,
	logger logger.Logger,
	userRepo user.UserRepository,
	walletRepo wallet.WalletRepository,
) *User {
	return &User{
		uow:        uow,
		jwt:        jwt,
		logger:     logger,
		userRepo:   userRepo,
		walletRepo: walletRepo,
	}
}

func (a *User) GetProfile(ctx context.Context) (*dto.ProfileResponse, error) {
	var (
		userId = user.UserId(ctx.Value("user_id").(string))
	)

	user, err := userId.GetUser(ctx, a.userRepo)
	if err != nil {
		return nil, err
	}

	return &dto.ProfileResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (a *User) Login(ctx context.Context, payload dto.LoginInput) (*dto.JWTToken, error) {
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

	return &dto.JWTToken{
		AccessToken:  token.AccessToken,
		RefreshToken: "",
	}, nil
}

func (a *User) Register(ctx context.Context, payload dto.RegisterInput) error {
	userId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	walletId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	now := time.Now()

	user := &user.User{
		ID:        userId.String(),
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: now,
		UpdatedAt: now,
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

	newWallet := wallet.Wallet{
		Id:        walletId.String(),
		UserID:    userId.String(),
		Amount:    0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = a.uow.Execute(ctx, func(ctx context.Context) (result *unitofwork.Result, err error) {
		err = newWallet.CreateNewWallet(ctx, a.walletRepo)
		if err != nil {
			return nil, err
		}

		err = user.NewUser(ctx, a.userRepo)
		if err != nil {
			return nil, err
		}
		return &unitofwork.Result{}, err
	})

	if err != nil {
		return err
	}

	return nil
}
