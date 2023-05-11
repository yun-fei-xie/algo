package _2_slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/max-consecutive-ones-iii/
给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。

这个题的精髓在于，不需要真的翻转，而是记录窗口中0的个数
当窗口中的0个个数<=k ，窗口都是合法的。于是再次变成简单的滑动窗口问题。
用一个map记录区间中有多少一个1，多少个0。

*/

func longestOnes(nums []int, k int) int {
	mp0 := make(map[int]int)
	mp0[1] = 0
	mp0[0] = 0                 // 多于
	maxLength := math.MinInt64 // 不需要
	left := 0                  // 可以放到for中
	right := 0
	for right < len(nums) {
		mp0[nums[right]]++
		if count, _ := mp0[0]; count <= k { // 窗口合法，最多包含k个0

			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}
			right++ // if else 同时出现 right++ 可以放到for表达式中

		} else { // 窗口不合法,从尾部拿掉一个元素
			for {
				c, _ := mp0[0]
				if c <= k {
					break
				} else {
					mp0[nums[left]]--
					left++
				}
			}
			right++
		}
	}
	if maxLength == math.MinInt64 {
		return 0
	}
	return maxLength
}

// 更加简洁的代码，上一个解写的荣誉主要是因为map用的不熟悉。
// map查找一个不存在的值会返回value类型的字面量。而且map可以直接用++ --
func longestOnes2(nums []int, k int) int {
	var ans int
	counter := make(map[int]int)

	for left, right := 0, 0; right < len(nums); right++ {
		counter[nums[right]]++
		if counter[0] <= k {
			ans = max(ans, right-left+1)
			continue
		} else {
			for counter[0] > k {
				counter[nums[left]]--
				left++
			}
		}
	}
	return ans
}

func TestLongestOnes(t *testing.T) {

	//fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes2([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))

}
