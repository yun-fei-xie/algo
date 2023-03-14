package _2_slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

/*
解题思路：滑动窗口
1. left代表窗口的左边，right代表窗口的右边。如果区间合法，则right++
2. right++后，判断区间是否合法，如果合法则更新最大长度。
3. 如果区间不合法，则left++ 直到区间合法（left++过程中，区间不会一直不合法（最差的情况下区间中剩下一个字符））
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := math.MinInt64
	mp := make(map[uint8]int, 0)
	left := 0
	right := 0
	for ; right < len(s); right++ {
		mp[s[right]]++
		if (right-left)+1 == len(mp) {
			if (right - left + 1) > max {
				max = right - left + 1
			}
		} else { // 当窗口的长度大于mp中key的个数，则代表有重复
			for right-left+1 > len(mp) { // 不断消除重复元素
				mp[s[left]]--
				if mp[s[left]] == 0 {
					delete(mp, s[left])
				}
				left++
			}
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
