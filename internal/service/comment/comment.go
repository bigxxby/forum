package comment

import "forum/internal/repository/comment"

type CommentService struct {
	CommentRepo *comment.CommentRepo
}

func NewCommentService(repo *comment.CommentRepo) *CommentService {
	return &CommentService{
		CommentRepo: repo,
	}
}
