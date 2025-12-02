package main

import (
	"fmt"

	"gorm.io/gorm"
)

func prepareData(db *gorm.DB) {
	user := User{Name: "张三"}
	db.Debug().Create(&user)

	post1 := Post{Title: "标题1", Content: "GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象", UserID: user.ID}
	db.Debug().Create(&post1)

	comments1 := []Comment{{Content: "这是第1条评论", PostID: post1.ID},
		{Content: "这是第2条评论", PostID: post1.ID},
		{Content: "这是第3条评论", PostID: post1.ID},
		{Content: "这是第4条评论", PostID: post1.ID},
		{Content: "这是第5条评论", PostID: post1.ID}}
	db.Debug().Create(&comments1)

	post2 := Post{Title: "标题2", Content: "当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误", UserID: user.ID}
	db.Debug().Create(&post2)

	comments2 := []Comment{{Content: "这是第11条评论", PostID: post2.ID},
		{Content: "这是第22条评论", PostID: post2.ID},
		{Content: "这是第33条评论", PostID: post2.ID}}
	db.Debug().Create(&comments2)
}

func findPostWithMostComments(db *gorm.DB) Post {
	var post Post
	db.Debug().Raw(`SELECT * FROM posts 
	ORDER BY (SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) DESC 
	LIMIT 1`).Scan(&post)
	fmt.Println(post)
	return post
}

func main() {
	db := InitDB()

	//插入数据
	//prepareData(db)

	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	// user := User{ID: 1}
	// db.Debug().Preload("Posts").Preload("Posts.Comments").First(&user, user.ID)
	// fmt.Println(user)

	//使用Gorm查询评论数量最多的文章信息
	findPostWithMostComments(db)
}
