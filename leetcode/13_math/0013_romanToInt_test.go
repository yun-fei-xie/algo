package _3_math

import (
	"fmt"
	"testing"
)

/*
罗马数字转整数：https://leetcode.cn/problems/roman-to-integer/
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func romanToInt(s string) int {
	var m = map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var ans int
	var pre int = 1
	for i := len(s) - 1; i >= 0; i-- {
		num := m[s[i]]
		if num >= pre {
			ans += num
			pre = num
		} else {
			ans -= num
		}
	}
	return ans
}

func TestRomanToInt(t *testing.T) {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
}
