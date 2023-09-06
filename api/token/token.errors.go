package token

import "errors"

var (
	ErrAlreadyExists = errors.New("token already exists")
	ErrNotFound      = errors.New("token not found")
	ErrInternal      = errors.New("internal error")
)
