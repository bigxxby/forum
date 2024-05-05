package models

type User struct {
	ID       int    `json:"id" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
