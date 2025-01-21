package domain

import "time"

type NotionEvent struct {
	Id       string    `json:"id,omitempty"`
	Url      string    `json:"url,omitempty"`
	RefId    string    `json:"refId,omitempty"`
	Name     string    `json:"name,omitempty"`
	Datetime time.Time `json:"datetime,omitempty"`
}
