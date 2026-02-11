package error

import errDomain "tlab/src/domain/error"

var ErrorMap map[error]ErrorAttributes = map[error]ErrorAttributes{
	errDomain.ErrInsufficientAmount: *NewErrorAttributes(400, "error", "insufficient amount to perform transfer."),
	errDomain.ErrUserNotFound:       *NewErrorAttributes(400, "error", "user not found, please check it again."),
	errDomain.ErrDuplicateEmail:     *NewErrorAttributes(400, "error", "duplicate email, try another email."),
	errDomain.ErrPasswordNotMatch:   *NewErrorAttributes(400, "error", "password not match, please check it again."),
	errDomain.ErrWalletNotFound:     *NewErrorAttributes(400, "error", "wallet not found, please check it again."),
	errDomain.ErrTransferError:      *NewErrorAttributes(500, "error", "something went wrong, try again later."),
}

type ErrorAttributes struct {
	Status  string
	Code    int
	Message string
}

func NewErrorAttributes(code int, status string, message string) *ErrorAttributes {
	return &ErrorAttributes{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
