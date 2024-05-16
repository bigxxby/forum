package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	UserId    int       `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Likes     int       `json:"likes"`
	Liked     bool      `json:"liked"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
