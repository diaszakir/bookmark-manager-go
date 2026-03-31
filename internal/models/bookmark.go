package models

type Bookmark struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
	URL  string `json:"url"`
	Tag  string `json:"tag"`
}
