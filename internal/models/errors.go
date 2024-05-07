package models

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInternalServer     = errors.New("internal server error")
	ErrConflict           = errors.New("conflict")
)
