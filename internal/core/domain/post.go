package domain

import "time"

type Post struct {
	Id   string    `json:"id"`
	Body string    `json:"body"`
	Date time.Time `json:"date"`
}
