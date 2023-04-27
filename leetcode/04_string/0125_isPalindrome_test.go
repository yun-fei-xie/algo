package _4_string

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/valid-palindrome/
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

指针移动的时候做点手脚
*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1

	for left <= right {
		for left <= right && !isAlnum(s[left]) {
			left++
		}
		for left <= right && !isAlnum(s[right]) {
			right--
		}
		if left <= right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}

	}
	return true

}

func isAlnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}
