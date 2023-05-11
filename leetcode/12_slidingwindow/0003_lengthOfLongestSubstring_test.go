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

func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := math.MinInt64
	mp := make(map[uint8]int, 0)
	left := 0
	right := 0
	for ; right < len(s); right++ {
		mp[s[right]]++
		// 用hash表记录区间中的元素。加入hash表中只有6个key,区间长度为8，那么区间中有2个重复的字符
		// 还有一种写法，就是让hash表中永远不会有重复的元素
		if (right-left)+1 == len(mp) {
			if (right - left + 1) > ans {
				ans = right - left + 1
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
	return ans
}

/*
方法2：思想一样，但是代码简洁许多。
如果right的加入让字符有重复，那么不断消除重复字符。
这样，开启下一轮外层循环的时候，map中[left,right]这个区间一定没有重复字符。
counter[s[left]]-- 操作相当于去掉了区间左边的字符。
abcabcbb
*/
func lengthOfLongestSubstring2(s string) int {
	counter := make(map[uint8]int)
	var ans int
	for right, left := 0, 0; right < len(s); right++ {
		counter[s[right]]++
		for counter[s[right]] > 1 {
			counter[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	res := lengthOfLongestSubstring2(s)
	fmt.Println(res)

	res2 := lengthOfLongestSubstring1("bbbbb")
	fmt.Println(res2)
	res3 := lengthOfLongestSubstring1("pwwkew")
	fmt.Println(res3)
}
