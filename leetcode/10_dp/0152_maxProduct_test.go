package _0_dp

import (
	"fmt"
	"testing"
)

/*
152. 乘积最大子数组
https://leetcode.cn/problems/maximum-product-subarray/

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
这个题和53号问题，和最大的子数组略微不同
由于存在负数，那么会导致最大的变最小的，最小的变最大的。因此还需要维护当前最小值imin。

牢记状态的定义，一定以下标 i 结尾，即：乘积数组中 nums[i] 必须被选取。
*/
func maxProduct(nums []int) int {
	length := len(nums)
	maxDp := make([]int, length)
	minDp := make([]int, length)
	maxDp[0] = nums[0]
	minDp[0] = nums[0]
	var ans = maxDp[0]
	for i := 1; i < length; i++ {

		if nums[i] >= 0 {
			maxDp[i] = max(maxDp[i-1]*nums[i], nums[i])
			minDp[i] = min(minDp[i-1]*nums[i], nums[i])
		} else {
			maxDp[i] = max(minDp[i-1]*nums[i], nums[i])
			minDp[i] = min(maxDp[i-1]*nums[i], nums[i])
		}
		ans = max(ans, maxDp[i])
	}
	return ans
}

func min(args ...int) int {
	m := args[0]
	for _, item := range args {
		if item < m {
			m = item
		}
	}
	return m
}

func TestMaxProduct(t *testing.T) {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-2}))
}
