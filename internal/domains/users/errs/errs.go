package errs

import "errors"

var (
	ErrUserWithChatIDAlreadyExists = errors.New("user with chat id already exists")
	ErrUserNotFound                = errors.New("user not found")

	ErrFieldRequiredForUpdate = errors.New("at least 1 field is required to update")
)
