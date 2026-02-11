package constants

import "golang.org/x/crypto/bcrypt"

type (
	Role           string
	MerchantStatus string
)

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"

	PasswordMinLength = 6
	BcryptCost        = bcrypt.DefaultCost

	MerchantStatusActive   MerchantStatus = "active"
	MerchantStatusInactive MerchantStatus = "inactive"
)
