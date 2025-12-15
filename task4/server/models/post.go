package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"metanode.com/homework/server/utils"
)

type Posts struct {
	gorm.Model
	Title   string `gorm:"size:100;not null" json:"title"` // 标题
	Content string `gorm:"type:longtext" json:"content"`   // 内容
	UserID  uint   `gorm:"not null" json:"userId"`         // 用户ID（外键）
}

func (p *Posts) AddPost(tx *gorm.DB) *utils.AppError {

	if len(p.Title) <= 0 || len(p.Content) <= 0 || p.UserID == 0 {
		return utils.WrapError(utils.ErrInvalidRequest, errors.New("title or content can not be empty"))
	}

	if err := tx.Create(&p).Error; err != nil {
		return utils.WrapError(utils.ErrDatabase, err)
	}

	return nil

}

func (p *Posts) UpdatePost(tx *gorm.DB) *utils.AppError {

	if len(p.Title) <= 0 || len(p.Content) <= 0 || p.ID == 0 {
		return utils.WrapError(utils.ErrInvalidRequest, errors.New("ID title content can not be empty"))
	}

	result := tx.Model(&Posts{}).Where("id = ?", p.ID).Where("user_id = ?", p.UserID).Updates(p)

	if result.Error != nil {
		return utils.WrapError(utils.ErrDatabase, errors.New("failed to update post"))
	}

	if result.RowsAffected == 0 {
		return utils.WrapError(utils.ErrNotFound, fmt.Errorf("post [id: %d] not found", p.ID))
	}

	return nil
}

func (p *Posts) GetPostByID(tx *gorm.DB) (*Posts, *utils.AppError) {

	if p.ID == 0 {
		return nil, utils.WrapError(utils.ErrInvalidRequest, errors.New("postID can not be empty"))
	}

	var post Posts
	if err := tx.First(&post, p.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.WrapError(utils.ErrNotFound, fmt.Errorf("post [id: %d] not found", p.ID))
		}
		return nil, utils.WrapError(utils.ErrDatabase, err) // 数据库错误
	}

	return &post, nil
}

func (p *Posts) DeletePost(tx *gorm.DB) *utils.AppError {
	//必须是用户删除自己的post才行
	result := tx.Where("id = ?", p.ID).Where("user_id = ?", p.UserID).Delete(&Posts{})

	if result.Error != nil {
		return utils.WrapError(utils.ErrDatabase, result.Error)
	}

	if result.RowsAffected == 0 {
		return utils.WrapError(utils.ErrNotFound, fmt.Errorf("post [id: %d] not found", p.ID))
	}

	return nil
}

func (p *Posts) GetPosts(tx *gorm.DB) ([]Posts, *utils.AppError) {
	var posts []Posts
	if err := tx.Where("user_id = ?", p.UserID).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, utils.WrapError(utils.ErrDatabase, err)
	}

	return posts, nil

}
