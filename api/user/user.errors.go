package user

import "errors"

var (
	ErrPermissionDenied = errors.New("Permission denied!")
	ErrAlreadyEsists    = errors.New("User already exists!")
	ErrNotFound         = errors.New("User not found!")
	ErrInternal         = errors.New("Internal error!")
)
