package main

import "fmt"

func add10(p *int) {
	*p += 10
}

func main() {
	a := 23
	fmt.Println("before add:", a)
	add10(&a)
	fmt.Println("after add:", a)
}
