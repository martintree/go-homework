package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employees struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float32 `db:"salary"`
}

// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
func findTechDeptEmployees(db *sqlx.DB) []Employees {
	var employees []Employees

	db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")

	return employees
}

// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func findHighestSalaryEmployee(db *sqlx.DB) Employees {
	var employee Employees

	db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")

	return employee
}
func main() {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	employees := findTechDeptEmployees(db)
	log.Println("技术部员工信息：")
	for _, employee := range employees {
		log.Println(employee.Name, employee.Department, employee.Salary)
	}

	employee := findHighestSalaryEmployee(db)
	log.Printf("工资最高员工信息：姓名：%s，部门：%s，工资：%.2f", employee.Name, employee.Department, employee.Salary)

}
