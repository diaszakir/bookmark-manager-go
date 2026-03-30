package models

type Bookmark struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Tag  string `json:"tag"`
}
