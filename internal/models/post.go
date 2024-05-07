package models

import "time"

type Post struct {
	ID        int
	UserId    int
	Name      string
	Content   string
	CreatedAt time.Time
}
