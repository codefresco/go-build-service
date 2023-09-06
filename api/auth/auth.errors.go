package auth

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrInternal         = errors.New("internal error")
)
