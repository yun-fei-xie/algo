package _4_string

import (
	"fmt"
	"testing"
)

/*
344. 反转字符串
https://leetcode.cn/problems/reverse-string/description/
*/
func reverseString(s []byte) {

	left := 0
	right := len(s) - 1

	for left < right {

		s[left], s[right] = s[right], s[left] // golang的语法糖
		left++
		right--
	}

}

func TestReverseString(t *testing.T) {

	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))

}
