package dto

import "metanode.com/homework/server/models"

type CommentRequest struct {
	ID      uint   `json:"id"`
	Content string `json:"content" binding:"required,min=1,max=10000"`
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
}

func ToCreateCommentModel(req *CommentRequest) *models.Comments {
	return &models.Comments{
		PostID:  req.PostID,
		Content: req.Content,
		UserID:  req.UserID,
	}
}
