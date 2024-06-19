package models

import "time"

type Post struct {
	ID         int       `json:"id"`
	UserId     int       `json:"userId"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Categories []string  `json:"categories"`
	Category   string    `json:"categoriy"` //////////////////
	Likes      int       `json:"likes"`
	Dislikes   int       `json:"dislikes"`
	Liked      bool      `json:"liked"`
	Disliked   bool      `json:"disliked"`
	CreatedBy  string    `json:"createdBy"`
	CreatedAt  time.Time `json:"createdAt"`
}
