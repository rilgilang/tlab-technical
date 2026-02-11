package error

import "errors"

var (
	ErrUserNotFound     error = errors.New("user_not_found")
	ErrDuplicateEmail   error = errors.New("email_already_used")
	ErrPasswordNotMatch       = errors.New("password_not_match")

	ErrWalletNotFound     error = errors.New("wallet_not_found")
	ErrInsufficientAmount error = errors.New("insufficient_amount")
	ErrTransferError            = errors.New("server_error")
)
