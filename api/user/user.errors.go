package user

import "errors"

var (
	ErrAlreadyExists = errors.New("user already exists")
	ErrNotFound      = errors.New("user not found")
	ErrInternal      = errors.New("internal error")
)
