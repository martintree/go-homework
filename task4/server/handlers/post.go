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

func AddPost(c *gin.Context) {
	var toAddPost dto.PostRequest
	if err := c.ShouldBindJSON(&toAddPost); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": "invalid params"})
		return
	}
	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	toAddPost.UserID = userID.(uint)

	//dto转model
	post := dto.ToCreatePostModel(&toAddPost)
	if err := post.AddPost(db.GetDB()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post, "error": ""})
}

func GetPost(c *gin.Context) {
	postIDStr := c.Param("id")
	if postIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": "id can not be empty"})
		return
	}

	postID, err := strconv.ParseInt(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "", "error": "can not parse id to uint",
		})
		return
	}

	var post = models.Posts{}
	post.ID = uint(postID)

	respPost, err := post.GetPostByID(db.GetDB())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": respPost, "error": ""})
}

func DeletePost(c *gin.Context) {
	postIDStr := c.Param("id")
	if postIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": "id can not be empty"})
		return
	}

	postID, err := strconv.ParseInt(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "", "error": "can not parse id to uint",
		})
		return
	}

	var post = models.Posts{}
	post.ID = uint(postID)
	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	post.UserID = userID.(uint)

	if err := post.DeletePost(db.GetDB()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}

	// 同时删除该post下的所有comment
	var comment = models.Comments{PostID: post.ID}
	if err := comment.DeleteComment(db.GetDB()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": postID, "error": ""})
}

func UpdatePost(c *gin.Context) {
	var toUpdatePost dto.PostRequest
	if err := c.ShouldBindJSON(&toUpdatePost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": "invalid params"})
		return
	}

	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	toUpdatePost.UserID = userID.(uint)

	post := dto.ToUpdatePostModel(&toUpdatePost)
	if err := post.UpdatePost(db.GetDB()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post.ID, "error": ""})
}

func GetPosts(c *gin.Context) {
	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	var post = models.Posts{UserID: userID.(uint)}
	posts, err := post.GetPosts(db.GetDB())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts, "error": ""})
}
