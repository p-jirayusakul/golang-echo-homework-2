package utils

import "errors"

var (
	ErrDataNotFound       = errors.New("data nof found")
	ErrEmailIsAlreadyUsed = errors.New("this email is already used")
	ErrLoginFail          = errors.New("username or password invalid")
)
