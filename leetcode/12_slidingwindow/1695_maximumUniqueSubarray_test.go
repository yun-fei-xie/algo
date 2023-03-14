package _2_slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

/*
给你一个正整数数组 nums ，请你从中删除一个含有 若干不同元素 的子数组。删除子数组的 得分 就是子数组各元素之 和 。
返回 只删除一个 子数组可获得的 最大得分 。
如果数组 b 是数组 a 的一个连续子序列，即如果它等于 a[l],a[l+1],...,a[r] ，那么它就是 a 的一个子数组。

这个题最重要的就是转换题意
转换题意就是：最大和的连续子数组（子数组中元素不能重复）这样就非常容易理解和写出代码
*/

func maximumUniqueSubarray(nums []int) int {
	mp := make(map[int]int) // 判断窗口合法性
	left := 0
	right := 0
	sum := 0
	max := math.MinInt64

	for left <= right && right < len(nums) {

		sum += nums[right] // 放进来就有可能重复
		mp[nums[right]]++
		if right-left+1 == len(mp) {
			if sum > max {
				max = sum
			}
		}

		for right-left+1 > len(mp) {
			mp[nums[left]]--
			if mp[nums[left]] == 0 {
				delete(mp, nums[left])
			}
			sum -= nums[left]
			left++
		}
		right++
	}

	return max
}

func TestMaxinumUniqueSubArr(t *testing.T) {
	nums := []int{5, 2, 1, 2, 5, 2, 1, 2, 5}
	res := maximumUniqueSubarray(nums)
	fmt.Println(res)
}
