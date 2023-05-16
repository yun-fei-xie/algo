package kanpsack

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/partition-equal-subset-sum/

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

方法：计算出数组的和sum（sum必须为偶数），然后用回溯找出有没有几个元素的和=sum/2   本质和494问题是同一个问题

解法2：选与不选的问题，本质是一个0-1背包问题

*/

func canPartition2(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}

	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum/2+1)
	}
	dp[0][0] = true

	for i := 1; i < len(dp); i++ {
		for c := 0; c < len(dp[0]); c++ {
			if c < nums[i-1] {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = dp[i-1][c] || dp[i-1][c-nums[i-1]]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	// 是否存在方案，使得从nums[0...i]中挑选一些数字，使得这些数字的和恰好等于c
	var dfs func(i int, c int) bool
	dfs = func(i int, c int) bool {
		if i < 0 {
			if c == 0 {
				return true
			}
			return false
		}
		// 加入剪枝 减少不必要的递归 nums[i]是正整数 所以 c-nums[i] < 0 就不要递归了
		if c < nums[i] {
			return dfs(i-1, c)
		} else {
			return dfs(i-1, c-nums[i]) || dfs(i-1, c)
		}
		//	return dfs(i-1, c-nums[i]) || dfs(i-1, c)
	}
	return dfs(len(nums)-1, sum/2)
}

func TestCanPartition(t *testing.T) {
	fmt.Println(canPartition([]int{1, 5, 11, 15}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
	fmt.Println(canPartition2([]int{1, 5, 11, 15}))
	fmt.Println(canPartition2([]int{1, 2, 3, 5}))
}
