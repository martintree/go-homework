package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	PostCount int
	Posts     []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID             int `gorm:"primaryKey;autoIncrement"`
	Title          string
	Content        string
	UserID         int
	CommentsStatus int
	Comments       []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Content string
	PostID  int
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:admin123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	//建表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	fmt.Println("建表成功")
	return db
}

// func main() {
// 	 InitDB()
// }
