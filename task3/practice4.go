package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

// 使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
func findBooks(db *sqlx.DB) []Book {
	var books []Book

	db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)

	return books
}

func main() {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	books := findBooks(db)
	for _, book := range books {
		log.Println(book.ID, book.Title, book.Author, book.Price)
	}
}
