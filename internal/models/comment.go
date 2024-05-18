package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	PostID    *int      `json:"postId"`
	UserID    int       `json:"userId"`
	Content   string    `json:"content"`
	Edited    bool      `json:"edited"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	Liked     bool      `json:"liked"`
	Disliked  bool      `json:"disliked"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
