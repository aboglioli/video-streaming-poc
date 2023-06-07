package domain

type VideoDTO struct {
	ID       string   `json:"id"`
	Original Source   `json:"original"`
	Parts    []Source `json:"parts"`
}

type SourceDTO struct {
	URL    string `json:"url"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}
