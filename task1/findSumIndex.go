package main

import "fmt"

func findSumIndex(nums []int, sumNum int) (idx1, idx2 int) {
	if len(nums) == 0 {
		idx1 = 0
		idx2 = 0
		return
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == sumNum {
				idx1 = i
				idx2 = j
				return idx1, idx2
			}
		}
	}

	return
}

func main() {
	nums := []int{2, 7, 8, 9}
	sumNum := 11
	index1, index2 := findSumIndex(nums, sumNum)
	fmt.Printf("slice=%v, sumNum=%d,find index: index1=%d, index2=%d\n", nums, sumNum, index1, index2)

	nums = []int{3, 3}
	sumNum = 6
	index1, index2 = findSumIndex(nums, sumNum)
	fmt.Printf("slice=%v, sumNum=%d,find index: index1=%d, index2=%d\n", nums, sumNum, index1, index2)
}
