package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	name string
}

type Circle struct {
	name string
}

func (r *Rectangle) Area() {
	fmt.Printf("invoke %s's Area()\n", r.name)
}

func (r *Rectangle) Perimeter() {
	fmt.Printf("invoke %s's Perimeter()\n", r.name)
}

func (c *Circle) Area() {
	fmt.Printf("invoke %s's Area()\n", c.name)
}

func (c *Circle) Perimeter() {
	fmt.Printf("invoke %s's Perimeter()\n", c.name)
}

func main() {
	rectangle := &Rectangle{"Rectangle"}
	circle := &Circle{"Circle"}

	rectangle.Area()
	rectangle.Perimeter()

	circle.Area()
	circle.Perimeter()
}
