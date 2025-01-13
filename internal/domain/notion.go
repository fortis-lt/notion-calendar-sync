package domain

import "time"

type NotionEvent struct {
	Id   string        `json:"id,omitempty"`
	Name string        `json:"name,omitempty"`
	Url  string        `json:"url,omitempty"`
	Date time.DateTime `json:"date,omitempty"`
}
