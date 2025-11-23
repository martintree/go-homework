package main

import (
	"fmt"
	"strconv"
)

func reverseString(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func isPalindromicNum(num int) bool {
	if num < 0 {
		return false
	}
	intStr := strconv.Itoa(num)
	return reverseString(intStr) == intStr

}

func main() {
	fmt.Println("6886 is palindromic number =", isPalindromicNum(6886))
	fmt.Println("123 is palindromic number =", isPalindromicNum(123))
	fmt.Println("1 is palindromic number =", isPalindromicNum(1))
}
