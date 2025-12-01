package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Students struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Name  string
	Age   int
	Grade string
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
func insertStudent(db *gorm.DB) {

	student := &Students{Name: "王五", Age: 19, Grade: "三年级"}
	result := db.Create(student)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("Student id=%d RowsAffected=%d", student.ID, result.RowsAffected)
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息
func findStudent(db *gorm.DB) {
	var student []Students

	db.Where("age > ?", 18).Find(&student)

	for _, student := range student {
		fmt.Printf("Student id=%d Name=%s Age=%d Grade=%s\n", student.ID, student.Name, student.Age, student.Grade)
	}
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func updateStudent(db *gorm.DB) {
	result := db.Debug().Model(&Students{}).Where("name = ?", "张三").Update("grade", "四年级")
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("RowsAffected=%d\n", result.RowsAffected)
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func deleteStudent(db *gorm.DB) {
	result := db.Debug().Where("age < ?", 15).Delete(&Students{})
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("RowsAffected=%d\n", result.RowsAffected)
}
func main() {
	db, err := gorm.Open(mysql.Open("root:admin123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	//建表
	db.AutoMigrate(&Students{})

	//插入数据
	//insertStudent(db)

	//查询数据
	//findStudent(db)

	//更新数据
	//updateStudent(db)

	deleteStudent(db)
}
