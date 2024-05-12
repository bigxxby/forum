package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	UserId    int       `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
