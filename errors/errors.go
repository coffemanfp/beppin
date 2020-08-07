package errors

import "errors"

// TODO: Create basic errors.
// TODO: Create database errors.
// TODO: Create server errors.

// Common errors.
var (
	ErrNotExistentObject          = errors.New("non-existent object")
	ErrExistentObject             = errors.New("existent object")
	ErrNotSupportedType           = errors.New("not-supported type")
	ErrNotProvidedOrInvalidObject = errors.New("not-provided or invalid object")
)

// System errors.
var (
	ErrInvalidSettings = errors.New("invalid settings")
)

// Server errors.
var (
	ErrInvalidBody      = "invalid body"
	ErrInvalidParam     = "invalid param"
	ErrInvalidUserLogin = "missing or invalid user credentials"
)
