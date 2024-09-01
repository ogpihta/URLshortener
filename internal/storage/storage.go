package storage

import "errors"

var (
	ErrNotFound  = errors.New("URL not found")
	ErrURLExists = errors.New("URL already exists")
)
