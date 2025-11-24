package main

import "fmt"

func plusOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		nums[i] = 0
	}
	//处理切片里都是9进位的情况
	nums = append([]int{1}, nums...)
	return nums
}

func main() {
	numbers := []int{1, 2, 3}
	fmt.Print(numbers)
	fmt.Println(" plus one result is:", plusOne(numbers))

	numbers = []int{4, 3, 2, 1}
	fmt.Print(numbers)
	fmt.Println(" plus one result is:", plusOne(numbers))

	numbers = []int{9}
	fmt.Print(numbers)
	fmt.Println(" plus one result is:", plusOne(numbers))
}
