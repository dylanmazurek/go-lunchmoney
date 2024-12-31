package lunchmoney

import "errors"

var (
	ErrAccessTokenExpired = errors.New("access token has expired")
	ErrRetryLimitReached  = errors.New("retry limit reached")
	ErrMissingCredentials = errors.New("missing credentials")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
