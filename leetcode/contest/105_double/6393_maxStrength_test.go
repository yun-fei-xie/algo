package _105_double_test

import (
	"fmt"
	"math"
	"testing"
)

/*
方法：暴力回溯（选或者不选）

方法:动态规划（因为有负数的情况，需要维护最大和最小）当考虑到nums[i]的时候，有四种转移情况。

	以最大值为例 maxDp[i] = max{nums[i] , nums[i]* maxDp[i-1] , dp[i-1] ,nums[i]*minDp[i-1] }
*/
func maxStrength(nums []int) int64 {
	var ans int64 = math.MinInt64
	var dfs func(i int, cnt int, strength int64)
	dfs = func(i int, cnt int, strength int64) {
		if i == len(nums) {
			if strength > ans && cnt != 0 {
				ans = strength
			}
			return
		}
		// 不选
		dfs(i+1, cnt, strength)
		// 选
		dfs(i+1, cnt+1, strength*int64(nums[i]))
	}

	dfs(0, 0, 1)
	return ans
}

func TestMaxStrength(t *testing.T) {
	fmt.Println(maxStrength([]int{3, -1, -5, 2, 5, -9}))
	fmt.Println(maxStrength([]int{0, -1}))
}
