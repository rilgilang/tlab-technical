package user

import (
	"context"
	"time"
	errorDomain "tlab/src/domain/error"
	"tlab/src/domain/sharedkernel/jwt"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserId string
	User   struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	RegisterInput struct {
		Name     string `json:"name" validate:"required,min=2,max=50"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	LoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	JWTToken struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func (u *User) NewUser(ctx context.Context, repository UserRepository) error {
	if err := repository.CreateUser(ctx, *u); err != nil {
		return err
	}
	return nil
}

func (u *LoginInput) GetUser(ctx context.Context, repository UserRepository) (*User, error) {
	return repository.GetUserByEmail(ctx, u.Email)
}

func (u *UserId) GetUser(ctx context.Context, repository UserRepository) (*User, error) {
	return repository.GetUserById(ctx, string(*u))
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errorDomain.ErrPasswordNotMatch
	}
	return nil
}

func (u *User) IsEmailDuplicate(ctx context.Context, repository UserRepository) (bool, error) {
	user, err := repository.GetUserByEmail(ctx, u.Email)
	if err != nil {
		return true, err
	}

	if user != nil && user.Email == u.Email {
		return true, errorDomain.ErrDuplicateEmail
	}

	return false, nil
}

func (u *User) GenerateSalt() ([]byte, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func (u *User) GenerateJWTToken(ctx context.Context, jwt jwt.JWT) (*JWTToken, error) {
	accessToken, err := jwt.GenerateAccessToken(ctx, u.ID)

	if err != nil {
		return nil, err
	}

	return &JWTToken{
		AccessToken:  *accessToken,
		RefreshToken: "",
	}, nil
}
