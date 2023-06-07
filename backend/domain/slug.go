package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidSlug = errors.New("invalid slug")
)

type Slug string

func NewSlug(id string) (Slug, error) {
	if len(id) < 6 {
		return "", fmt.Errorf("%w: %s", ErrInvalidSlug, id)
	}

	return Slug(id), nil
}

func (id Slug) String() string {
	return string(id)
}
