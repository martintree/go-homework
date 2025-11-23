package main

import "fmt"

func findOnceNumber(nums [11]int) int {
	count := make(map[int]int)

	//遍历数组依靠map的key相同的会加一的算法来找出只出现一次的整数
	for _, num := range nums {
		count[num]++
	}

	for numKey, cnt := range count {
		if cnt == 1 {
			return numKey
		}
	}
	return 0
}

func main() {
	arry := [...]int{1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6}

	fmt.Println(findOnceNumber(arry))
}
