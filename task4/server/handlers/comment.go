package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/db"
	"metanode.com/homework/server/dto"
	"metanode.com/homework/server/models"
)

func AddComment(c *gin.Context) {
	var toAddComment dto.CommentRequest
	if err := c.ShouldBindJSON(&toAddComment); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": "invalid params"})
		return
	}

	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	toAddComment.UserID = userID.(uint)
	//dto转model
	comment := dto.ToCreateCommentModel(&toAddComment)
	if err := comment.AddComment(db.GetDB()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment, "error": ""})
}

func DeleteComment(c *gin.Context) {
	postIDStr := c.Param("postId")
	if postIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": "post id can not be empty"})
		return
	}

	postID, err := strconv.ParseInt(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "", "error": "can not parse id to uint",
		})
		return
	}

	var comment = models.Comments{PostID: uint(postID)}
	if err := comment.DeleteComment(db.GetDB()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": postID, "error": ""})
}
