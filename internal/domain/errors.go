package domain

import (
	"errors"
)

// ErrNotFound not found
var ErrNotFound = errors.New("Not found")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot be deleted")

// ErrInvalidUsername invalid username
var ErrInvalidUsername = errors.New("Invalid username")

// ErrDuplicatedUsername duplicated username
var ErrDuplicatedUsername = errors.New("Username already exist")

// ErrInvalidEmail invalid email
var ErrInvalidEmail = errors.New("Invalid email")

// ErrInvalidPassword invalid password
var ErrInvalidPassword = errors.New("Invalid password")
