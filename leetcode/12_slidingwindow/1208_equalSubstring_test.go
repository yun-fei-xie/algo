package _2_slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/get-equal-substrings-within-budget/description/

滑动窗口->不定窗口

给你两个长度相同的字符串，s 和 t。
将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。
用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。
如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。
如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。
*/

func equalSubstring(s string, t string, maxCost int) int {

	var abs func(a, b uint8) int
	abs = func(a, b uint8) int {
		if a > b {
			return int(a - b)
		}
		return int(b - a)
	}

	maxLength := math.MinInt64
	left := 0
	right := 0

	for right < len(s) {

		cost := abs(s[right], t[right])

		maxCost = maxCost - cost

		if maxCost >= 0 {
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}

		} else { // maxCost <0
			for maxCost < 0 {
				maxCost += abs(s[left], t[left])
				left++
			}
		}
		right++
	}
	if maxLength == math.MinInt64 {
		return 0
	}
	return maxLength
}

func TestEqualSubString(t *testing.T) {
	res := equalSubstring("abcd", "bcdf", 3)
	fmt.Println(res)
}
