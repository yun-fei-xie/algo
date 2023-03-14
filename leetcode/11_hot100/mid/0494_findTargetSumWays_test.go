package mid

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/target-sum/?favorite=2cktkvj

给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。


最简单的方法就是枚举

应该还有动态规划的方法

*/

func findTargetSumWays(nums []int, target int) int {
	var res int

	var dfs func(arr []int, startIndex int, sum int)
	dfs = func(arr []int, startIndex int, sum int) {
		if startIndex >= len(arr) {
			if sum == target {
				res++
			}
			return
		}

		dfs(arr, startIndex+1, sum+arr[startIndex])
		dfs(arr, startIndex+1, sum-arr[startIndex])
	}

	dfs(nums, 0, 0)
	return res
}

func TestFindTargetSumWays(t *testing.T) {
	res := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	fmt.Println(res)
}
