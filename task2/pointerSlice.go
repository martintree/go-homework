package main

import "fmt"

func mul2(p *[]int) {
	for i := range *p {
		(*p)[i] *= 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("before:", nums)
	mul2(&nums)
	fmt.Println("after:", nums)
}
