package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	Edited    bool      `json:"edited"`
	Likes     int       `json:"likes"`
	Liked     bool      `json:"liked"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"created_at"`
}
