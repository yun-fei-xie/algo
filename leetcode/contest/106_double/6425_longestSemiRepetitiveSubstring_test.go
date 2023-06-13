package _06_double

import (
	"fmt"
	"testing"
)

/*
52233

范围不大 尝试暴力枚举区间[i,j]

方法2：双指针
*/
func longestSemiRepetitiveSubstring(s string) int {

	var ans int
	length := len(s)
	if length == 1 {
		return 1
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if valid(&s, i, j) {
				if j-i+1 > ans {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}
func valid(s *string, i int, j int) bool {
	repet := false
	for k := i + 1; k <= j; {
		if (*s)[k] != (*s)[k-1] {
			k++
		} else {
			if repet == false {
				repet = true
			} else {
				return false
			}
			k++
		}
	}
	return true
}

func TestLongestSemi(t *testing.T) {
	fmt.Println(longestSemiRepetitiveSubstring("52233"))
	fmt.Println(longestSemiRepetitiveSubstring("5494"))
	fmt.Println(longestSemiRepetitiveSubstring("11111"))
}
