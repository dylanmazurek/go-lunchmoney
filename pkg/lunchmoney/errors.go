package lunchmoney

import "errors"

var (
	ErrAPIKeyEmpty        = errors.New("api key cannot be empty")
	ErrAccessTokenExpired = errors.New("access token has expired")
	ErrRetryLimitReached  = errors.New("retry limit reached")
	ErrMissingCredentials = errors.New("missing credentials")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
