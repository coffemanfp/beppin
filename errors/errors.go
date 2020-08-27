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
	ErrInvalidData                = errors.New("invalid data")
)

// System errors.
var (
	ErrInvalidSettings = errors.New("missing or invalid settings")
	ErrClosedDatabase  = errors.New("missing or closed database connetion")
)

// Client and server errors.
var (
	ErrInvalidBody       = "missing or invalid body"
	ErrInvalidParam      = "missing or invalid param"
	ErrInvalidUserLogin  = "missing or invalid user credentials"
	ErrInvalidAvatarData = "missing or invalid avatar data"
)
