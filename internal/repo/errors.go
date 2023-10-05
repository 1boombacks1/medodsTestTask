package repo

import "errors"

var (
	ErrAlreadyExists = errors.New("session already exists")
	ErrNotFound      = errors.New("session not found")
)
