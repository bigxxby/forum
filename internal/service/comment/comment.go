package comment

import (
	"forum/internal/repository/comment"
	"forum/internal/repository/likes"
)

type CommentService struct {
	CommentRepo *comment.CommentRepo
	LikeRepo    *likes.LikesRepo
}

func NewCommentService(repo *comment.CommentRepo, likes *likes.LikesRepo) *CommentService {
	return &CommentService{
		CommentRepo: repo,
		LikeRepo:    likes,
	}
}
