package models

import (
	"errors"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Content string `gorm:"type:longtext" json:"content"` // 内容
	UserID  uint   `gorm:"not null" json:"userId"`       // 用户ID（外键）
	PostID  uint   `gorm:"not null" json:"postId"`
}

func (c *Comments) AddComment(tx *gorm.DB) error {

	if len(c.Content) <= 0 || c.UserID == 0 || c.PostID == 0 {
		return errors.New("userId postId content can not be empty")
	}

	if err := tx.Create(&c).Error; err != nil {
		return errors.New("failed to create comment")
	}

	return nil
}

func (c *Comments) DeleteComment(tx *gorm.DB) error {

	result := tx.Where("post_id = ?", c.PostID).Delete(&Comments{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
