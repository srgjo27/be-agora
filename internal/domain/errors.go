package domain

import "errors"

var (
	ErrConflict = errors.New("data sudah ada")
	ErrNotFound = errors.New("data tidak ditemukan")
	ErrInvalid  = errors.New("data tidak valid")
)
