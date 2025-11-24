package main

import "fmt"

type Result struct {
	count       int
	resultSlice []int
}

func removeDuplicate(nums []int) Result {
	if len(nums) == 0 {
		return Result{0, []int{}}
	}

	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return Result{slow + 1, nums[:slow+1]}

}
func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 5, 5}
	result := removeDuplicate(nums)
	fmt.Println("count:", result.count, " result:", result.resultSlice)
}
