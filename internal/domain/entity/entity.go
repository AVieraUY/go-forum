package entity

import (
	"github.com/google/uuid"
)

// ID for entities
type ID = uuid.UUID

// NewID generates new ID
func NewID() ID {
	return ID(uuid.New())
}

// StringToID takes a string and parse it
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
