package authmodel

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidCode  = errors.New("invalid code")
)
