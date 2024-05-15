package comment

import "forum/internal/service/comment"

type CommentController struct {
	CommentService *comment.CommentService
}

func NewCommentController(service *comment.CommentService) *CommentController {
	return &CommentController{
		CommentService: service,
	}
}
