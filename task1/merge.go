package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果数组，第一个区间直接加入
	firstElmt := [][]int{intervals[0]}

	// 从第二个区间开始遍历
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := firstElmt[len(firstElmt)-1]

		// 检查是否重叠：当前区间的开始 <= 已合并区间的结束
		if current[0] <= lastMerged[1] {
			// 重叠，合并区间：更新结束位置为两者的最大值
			lastMerged[1] = max(lastMerged[1], current[1])
		} else {
			// 不重叠，直接添加到结果中
			firstElmt = append(firstElmt, current)
		}
	}

	return firstElmt
}

// 辅助函数：返回两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := [][]int{{6, 10}, {1, 3}, {2, 6}, {15, 18}}
	//先排序变成：{1, 3}, {2, 6}, {6, 10}, {15, 18}
	//再拿已经合并的数组的最后一个元素跟当前数组的第一个元素比较，如果大于等于则合并
	// （即将当前数组最后一个元素的值替换已合并数组最后一个元素的值）
	result := merge(nums)
	fmt.Println(result)
}
