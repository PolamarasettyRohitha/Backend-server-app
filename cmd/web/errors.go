package main

import "errors"

var (
	ErrMissingArguments        = errors.New("missing arguments")
	ErrInalidArguments         = errors.New("invalid arguments")
	ErrNotFound                = errors.New("not found")
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrAuthHeaderRequired      = errors.New("authorization header is required")
	ErrInvalidOrExpiredToken   = errors.New("invalid or expired token")
	ErrUnAuthorizedUser        = errors.New("unauthorized user")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidUser             = errors.New("invlaid user")
)
