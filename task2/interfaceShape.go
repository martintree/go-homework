package main

import (
	"fmt"
)

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

type Triangle struct {
	name string
}

// 注意这里是值传参
func (r Triangle) Area() {
	fmt.Printf("invoke %s's Area()\n", r.name)
}

func (r Triangle) Perimeter() {
	fmt.Printf("invoke %s's Perimeter()\n", r.name)
}

// 注意这里是指针传参
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
	circle := Circle{"Circle"}
	triangle := Triangle{"Triangle"}

	//语法糖，自动取地址,实际上是：(&triangle)
	(&triangle).Area()

	var s Shape = (&triangle)
	s.Area()

	fmt.Println("=========")
	//语法糖， 自动解引用，实际上是：(*rectangle)所以下面这句也可以
	(*rectangle).Area()
	rectangle.Perimeter()

	circle.Area()
	circle.Perimeter()

	var shapes []Shape
	shapes = append(shapes, rectangle)
	shapes = append(shapes, &circle)
	shapes = append(shapes, triangle)
	fmt.Println("=========")
	for _, shape := range shapes {
		shape.Area()
		shape.Perimeter()
	}
}
