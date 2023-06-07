package domain

import (
	"encoding/json"
)

type Source struct {
	URL    Slug
	Width  int32
	Height int32
}

func (s *Source) MarshalJSON() ([]byte, error) {
	source := &SourceDTO{
		URL:    s.URL.String(),
		Width:  s.Width,
		Height: s.Height,
	}

	return json.Marshal(source)
}
