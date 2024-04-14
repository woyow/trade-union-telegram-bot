package errs

import "errors"

var (
	ErrChatCurrentStateNotExists     = errors.New("chat current state not exists")
	ErrChatCurrentStateAlreadyExists = errors.New("chat current state already exists")

	ErrUnknownAnswer = errors.New("unknown answer")

	ErrStatusCodeUnsuccessful = errors.New("status code unsuccessful")
)
