package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
题目链接:https://leetcode.cn/problems/house-robber/
*/
/*
递归+记忆化搜索
*/
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	var dfs func(nums []int, startIndex int) int
	dfs = func(nums []int, startIndex int) int {
		if startIndex >= len(nums) { // 没房子可以打劫 直接返回
			return 0
		}
		if memo[startIndex] != -1 {
			return memo[startIndex]
		}

		// [startIndex , n-1] 扫荡这个区间 目的是求解打劫这个区间能够获得的最大收益
		max := math.MinInt64
		for i := startIndex; i < len(nums); i++ {
			m := nums[i] + dfs(nums, i+2)
			if m > max {
				max = m
			}
		}
		memo[startIndex] = max

		return memo[startIndex]
	}

	return dfs(nums, 0)
}

/*
dp解法：自底向上进行推导，先从最后开始推
*/
func robdp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	memo := make([]int, len(nums))
	memo[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 1; i >= 0; i-- { // 每一个节点
		// [i...n-1]
		max := math.MinInt64
		for j := i; j < len(nums); j++ {
			var m int
			if j+2 >= len(nums) { // 如果越界-> 对应递归到底 到达一个空的位置 此时直接返回0
				m = nums[j] + 0
			} else {
				m = nums[j] + memo[j+2]
			}
			if m > max {
				max = m
			}
		}
		memo[i] = max
	}
	return memo[0]
}

func TestRob(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
	fmt.Println(robdp(nums))
}
