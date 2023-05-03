package _0_dp

import (
	"fmt"
	"testing"
)

/*
这道题是第53号问题，最大连续子数组的和的一个变形问题。
放宽了条件，允许数组转起来。
*/

func maxSubarraySumCircular(nums []int) int {

	var sum = nums[0]
	var dp = make([]int, len(nums))
	var maxSub = nums[0]
	dp[0] = nums[0]
	// 求出不含环的最大子序列和 maxSub
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSub = max(maxSub, dp[i])
	}

	// [1...length-2]
	var minSub = 0
	var minDp = make([]int, len(nums))
	minDp[0] = 0
	for j := 1; j < len(nums)-1; j++ {
		minDp[j] = min(minDp[j-1]+nums[j], nums[j])
		minSub = min(minSub, minDp[j])
	}
	return max(maxSub, sum-minSub)
}

func TestMaxSubArraySumCircular(t *testing.T) {
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
}
