package main

import "fmt"

func isValidatedBracket(str string) bool {
	var stack []rune

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, chr := range str {
		switch chr {
		case '(', '[', '{':
			stack = append(stack, chr)
		case ')', ']', '}':
			//栈内为空说明没有左括号
			if len(stack) == 0 {
				return false
			}
			//取栈最顶端的元素
			topElm := stack[len(stack)-1]
			//将最顶端的弹出栈
			stack = stack[:len(stack)-1]
			if topElm != pairs[chr] {
				return false
			}
		}
	}

	// 最后栈必须为空
	return len(stack) == 0
}

func main() {
	str := "()"
	fmt.Println(str, "is validate bracket=", isValidatedBracket(str))

	str = "()[]{}"
	fmt.Println(str, "is validate bracket=", isValidatedBracket(str))

	str = "(]"
	fmt.Println(str, "is validate bracket=", isValidatedBracket(str))

	str = "([])"
	fmt.Println(str, "is validate bracket=", isValidatedBracket(str))

	str = "([)]"
	fmt.Println(str, "is validate bracket=", isValidatedBracket(str))

}
