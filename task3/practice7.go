package main

import (
	"gorm.io/gorm"
)

func (p *Post) BeforeCreate(ctx *gorm.DB) (err error) {

	if err := ctx.Model(User{}).Where("id = ? ", p.UserID).Update("post_count", gorm.Expr("post_count+ ?", 1)).Error; err != nil {
		return err
	}

	return
}

func (c *Comment) AfterDelete(ctx *gorm.DB) (err error) {
	var count int64

	ctx.Model(&Comment{}).Where("post_id = ? ", c.PostID).Count(&count)
	if count <= 0 {
		//comments_status=0表示 "无评论"
		if err := ctx.Model(Post{}).Where("id = ? ", c.PostID).Update("comments_status", 0).Error; err != nil {
			return err
		}
	}

	return
}

func addPost(db *gorm.DB, userId int) {
	post := Post{Title: "标题555", Content: "免费提供的这项服务可在简体中文和其他100 多种语言之间即时翻译字词、短语和网页", UserID: userId}
	db.Debug().Create(&post)
}

func main() {
	db := InitDB()
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
	//addPost(db, 1)

	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	db.Debug().Delete(&Comment{PostID: 4}, 7)

}
