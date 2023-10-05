package services

import "errors"

var (
	ErrInvalidRefreshSession = errors.New("invalid refresh session")
	ErrRefreshTokenExpired   = errors.New("refresh token expired")
)
