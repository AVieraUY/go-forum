package domain

import (
	"errors"
)

// ErrNotFound not found
var ErrNotFound = errors.New("Not found")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot be deleted")