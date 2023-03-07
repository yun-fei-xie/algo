package _5_doublePointer

import (
	"fmt"
	"math"
	"testing"
)

/*
解题思路：滑动窗口
如何对子串进行判重复->set (当窗口的长度大于map中元素的个数的时候，窗口不合法)
*/

func lengthOfLongestSubstring(s string) int {

	max := math.MinInt64
	set := make(map[uint8]int)
	i := 0
	j := 0
	for i <= j && j < len(s) {
		// 考察j
		if _, found := set[s[j]]; !found {
			set[s[j]] = 1

			l := j - i + 1
			if l > max {
				max = l
			}
			j++
		} else {
			set[s[j]]++

			set[s[i]]--
			if set[s[i]] == 0 {
				delete(set, s[i])
			}
			i++
		}
	}
	return max
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)

	res2 := lengthOfLongestSubstring("bbbbb")
	fmt.Println(res2)
	res3 := lengthOfLongestSubstring("pwwkew")
	fmt.Println(res3)
}
