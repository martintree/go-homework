package main

import "fmt"

func longestPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	//以第一个字符串未基准，遍历第一个字符串的每个字符去跟其他字符串的字符进行对比
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		// 检查其他字符串在相同位置是否具有相同字符
		for j := 1; j < len(strs); j++ {
			//i >= len(strs[j])是为了判断遍历后面的字符串长度不够已确定的前缀
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	//如果所有字符串都相同，则返回第一个字符串
	return strs[0]
}
func main() {

	strs := []string{"tree", "treat", "trend"}

	fmt.Println(longestPrefix(strs))

}
