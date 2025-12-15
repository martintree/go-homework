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

func AddPost(c *gin.Context) {
	var toAddPost dto.PostRequest
	if err := c.ShouldBindJSON(&toAddPost); err != nil {
		c.Error(err)
		return
	}
	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	toAddPost.UserID = userID.(uint)

	//dto转model
	post := dto.ToCreatePostModel(&toAddPost)
	if err := post.AddPost(db.GetDB()); err != nil {
		c.Error(err)
		return
	}
	utils.Success(c, post)
}

func GetPost(c *gin.Context) {
	postIDStr := c.Param("id")
	if postIDStr == "" {
		c.Error(errors.New("id can not be empty"))
		return
	}

	postID, err := strconv.ParseInt(postIDStr, 10, 32)
	if err != nil {
		c.Error(errors.New("can not parse id to uint"))
		return
	}

	var post = models.Posts{}
	post.ID = uint(postID)

	respPost, selectErr := post.GetPostByID(db.GetDB())
	if selectErr != nil {
		c.Error(selectErr)
		return
	}
	utils.Success(c, respPost)
}

func DeletePost(c *gin.Context) {
	postIDStr := c.Param("id")
	if postIDStr == "" {
		c.Error(errors.New("id can not be empty"))
		return
	}

	postID, err := strconv.ParseInt(postIDStr, 10, 32)
	if err != nil {
		c.Error(errors.New("can not parse id to uint"))
		return
	}

	var post = models.Posts{}
	post.ID = uint(postID)
	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	post.UserID = userID.(uint)

	if err := post.DeletePost(db.GetDB()); err != nil {
		c.Error(err)
		return
	}

	// 同时删除该post下的所有comment
	var comment = models.Comments{PostID: post.ID}
	if err := comment.DeletePostComment(db.GetDB()); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, postID)
}

func UpdatePost(c *gin.Context) {
	var toUpdatePost dto.PostRequest
	if err := c.ShouldBindJSON(&toUpdatePost); err != nil {
		c.Error(err)
		return
	}

	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	toUpdatePost.UserID = userID.(uint)

	post := dto.ToUpdatePostModel(&toUpdatePost)
	if err := post.UpdatePost(db.GetDB()); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, post.ID)
}

func GetPosts(c *gin.Context) {
	//从上下文中获取userId
	userID, _ := c.Get("userId") // 已通过中间件验证
	var post = models.Posts{UserID: userID.(uint)}
	posts, err := post.GetPosts(db.GetDB())
	if err != nil {
		c.Error(err)
		return
	}
	utils.Success(c, posts)
}
