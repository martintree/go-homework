package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/db"
	"metanode.com/homework/server/dto"
	"metanode.com/homework/server/models"
	"metanode.com/homework/server/utils"
)

func AddComment(c *gin.Context) {
	var toAddComment dto.CommentRequest
	if err := c.ShouldBindJSON(&toAddComment); err != nil {
		c.Error(err)
		return
	}

	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	toAddComment.UserID = userID.(uint)
	//dto转model
	comment := dto.ToCreateCommentModel(&toAddComment)
	if err := comment.AddComment(db.GetDB()); err != nil {
		c.Error(err)
		return
	}
	utils.Success(c, comment)
}

func DeletePostComment(c *gin.Context) {
	postIDStr := c.Param("postId")
	if postIDStr == "" {
		c.Error(errors.New("post id can not be empty"))
		return
	}

	postID, err := strconv.ParseInt(postIDStr, 10, 32)
	if err != nil {
		c.Error(errors.New("can not parse postId to uint"))
		return
	}

	var comment = models.Comments{PostID: uint(postID)}
	if err := comment.DeletePostComment(db.GetDB()); err != nil {
		c.Error(err)
		return
	}
	utils.Success(c, postID)
}

func DeleteComment(c *gin.Context) {
	IDStr := c.Param("id")
	if IDStr == "" {
		c.Error(errors.New("id can not be empty"))
		return
	}

	id, err := strconv.ParseInt(IDStr, 10, 32)
	if err != nil {
		c.Error(errors.New("can not parse id to uint"))
		return
	}

	var comment = models.Comments{}
	comment.ID = uint(id)
	if err := comment.DeleteComment(db.GetDB()); err != nil {
		c.Error(err)
		return
	}
	utils.Success(c, comment.ID)
}
