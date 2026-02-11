package user

import (
	"context"
	"time"
	errorDomain "tlab/src/domain/error"
	"tlab/src/domain/sharedkernel/constants"
	"tlab/src/domain/sharedkernel/jwt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      constants.Role `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// DTO: Register input
type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// DTO: Login input
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type JWTToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Factory: Create new user from register input
func (u *User) NewUser(ctx context.Context, repository UserRepository) error {
	if err := repository.CreateUser(ctx, *u); err != nil {
		return err
	}
	return nil
}

// Password setter with hashing
func (u *User) SetPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), constants.BcryptCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *LoginInput) GetUser(ctx context.Context, repository UserRepository) (*User, error) {
	return repository.GetUserByEmail(ctx, u.Email)
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
	accessToken, err := jwt.GenerateAccessToken(ctx, u.ID.String())

	if err != nil {
		return nil, err
	}

	return &JWTToken{
		AccessToken:  *accessToken,
		RefreshToken: "",
	}, nil
}
