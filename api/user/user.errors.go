package user

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrAlreadyEsists    = errors.New("user already exists")
	ErrNotFound         = errors.New("user not found")
	ErrInternal         = errors.New("internal error")
)
