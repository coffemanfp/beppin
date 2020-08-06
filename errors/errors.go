package errors

import "errors"

// TODO: Create basic errors.
// TODO: Create database errors.
// TODO: Create server errors.

// Common errors
var (
	ErrNotExistentObject = errors.New("non-existent object")
	ErrExistentObject    = errors.New("existent object")
	ErrNotSupportedType  = errors.New("not-supported type")

	ErrInvalidSettings = errors.New("invalid settings")
)
