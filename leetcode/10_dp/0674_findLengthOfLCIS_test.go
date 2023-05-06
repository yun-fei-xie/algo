package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。

这个问题和最长上升子序列的区别在于，这个问题要求子序列连续。
*/
func findLengthOfLCIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length) // dp[i] = 以nums[i]为结尾的最长连续递增序列的长度
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	var ans int = 1
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func findLengthOfLCIS2(nums []int) int {
	var maxLen = 1
	var dfs func(startIndex int) int
	dfs = func(startIndex int) int {
		// 递归到底只有一个元素，表示以startIndex开始的子数组的最长连续自增子序列的长度只有1
		if startIndex == len(nums)-1 {
			return 1
		}
		// 应该放在这里,确保递归一定能下去
		next := dfs(startIndex + 1)

		if nums[startIndex] < nums[startIndex+1] {
			//next := dfs(startIndex + 1)
			maxLen = max(maxLen, 1+next)
			return 1 + next

		} else {
			return 1
		}
	}
	dfs(0)
	return maxLen
}

func TestFindLengthOfLCIS(t *testing.T) {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 2, 7}))
	fmt.Println(findLengthOfLCIS2([]int{1, 3, 5, 2, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2, 2}))
	fmt.Println(findLengthOfLCIS2([]int{2, 2, 2, 2, 2, 2}))
	fmt.Println(findLengthOfLCIS2([]int{1, 3, 5, 4, 2, 3, 4, 5}))
}
