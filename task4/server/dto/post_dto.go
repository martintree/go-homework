package dto

import "metanode.com/homework/server/models"

type PostRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title" binding:"required,min=1,max=100"`
	Content string `json:"content" binding:"required,min=1,max=10000"`
	UserID  uint   `json:"userId"`
}

func ToCreatePostModel(req *PostRequest) *models.Posts {
	return &models.Posts{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}
}

func ToUpdatePostModel(req *PostRequest) *models.Posts {
	post := models.Posts{}
	post.ID = req.ID
	post.Content = req.Content
	post.Title = req.Title
	post.UserID = req.UserID
	return &post
}
