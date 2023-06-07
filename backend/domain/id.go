package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrInvalidID = errors.New("invalid id")
)

type ID string

func NewID(id string) (ID, error) {
	if len(id) < 6 {
		return "", fmt.Errorf("%w: %s", ErrInvalidID, id)
	}

	return ID(id), nil
}

func GenerateUUID() ID {
	return ID(uuid.New().String())
}

func (id ID) String() string {
	return string(id)
}
