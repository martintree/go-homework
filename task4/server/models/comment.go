package models

import (
	"errors"

	"gorm.io/gorm"
	"metanode.com/homework/server/utils"
)

type Comments struct {
	gorm.Model
	Content string `gorm:"type:longtext" json:"content"` // 内容
	UserID  uint   `gorm:"not null" json:"userId"`       // 用户ID（外键）
	PostID  uint   `gorm:"not null" json:"postId"`
}

func (c *Comments) AddComment(tx *gorm.DB) *utils.AppError {

	if len(c.Content) <= 0 || c.UserID == 0 || c.PostID == 0 {
		return utils.WrapError(utils.ErrInvalidRequest, errors.New("userId postId content can not be empty"))
	}

	if err := tx.Create(&c).Error; err != nil {
		return utils.WrapError(utils.ErrDatabase, errors.New("failed to create comment"))
	}

	return nil
}

func (c *Comments) DeletePostComment(tx *gorm.DB) *utils.AppError {

	result := tx.Where("post_id = ?", c.PostID).Delete(&Comments{})

	if result.Error != nil {
		return utils.WrapError(utils.ErrDatabase, result.Error)
	}

	return nil
}

func (c *Comments) DeleteComment(tx *gorm.DB) *utils.AppError {

	result := tx.Where("id = ?", c.ID).Delete(&Comments{})

	if result.Error != nil {
		return utils.WrapError(utils.ErrDatabase, result.Error)
	}

	return nil
}
