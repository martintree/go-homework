package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	P          Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("print employee info:%+v\n", e)
}

func main() {
	employee := &Employee{P:Person{Name: "foo", Age: 25}, EmployeeID: "123456"}
	employee.PrintInfo()
	employee.P.Name = "bar"
	employee.PrintInfo()
}
