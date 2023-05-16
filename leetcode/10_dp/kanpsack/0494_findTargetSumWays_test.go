package kanpsack

import (
	"fmt"
	"testing"
)

/*
494. 目标和
https://leetcode.cn/problems/target-sum/?favorite=2cktkvj

给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000


方法：0-1背包问题
这个题的本质是在一些数字的前面放+号，一些数字前面放负号。
假设放正号的数字和为p，数组的和是s,则被设置为负号的那些数字的和为s-p。
正负相加等于target，则有 p + (-(s-p)) = t
展开得到 p = (s+t)/2。
也就是说，从数组中找一组数字，它们的和等于(s+t)/2。
从这里可以看出来（s + t）必须要为偶数。
给出有多少组的数字满足需求。

递归方法：


动态规划：



*/

func findTargetSumWays1(nums []int, target int) int {
	var res int

	var dfs func(arr []int, startIndex int, sum int)
	dfs = func(arr []int, startIndex int, sum int) {
		if startIndex >= len(arr) {
			if sum == target {
				res++
			}
			return
		}
		// 纯暴力
		dfs(arr, startIndex+1, sum+arr[startIndex])
		dfs(arr, startIndex+1, sum-arr[startIndex])
	}

	dfs(nums, 0, 0)
	return res
}

/*
背包思想
*/
func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	target = target + sum
	if target%2 != 0 {
		return 0
	}
	target = target / 2
	// target表示背包的容量
	// dfs函数表示: 在nums[0...i]这个区间挑选一些数据，他们恰好装满背包。
	// i表示考虑当前第i件物品，c表示当前背包剩余的容量。
	// 每一个数字都可以选择或者不选择
	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		// if c < nums[i] -> dp[i][c] = dp[i-1][c]
		if c < nums[i] {
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i])
	}
	// 也是倒着思考 给前面留多少空间
	return dfs(len(nums)-1, target)
}

/*
翻译成动态规划
*/
func findTargetSumWays3(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	target = target + sum
	if target%2 != 0 || target < 0 {
		return 0
	}
	target = target / 2
	// 整体平移一个单位的位移
	dp := make([][]int, len(nums)+1)
	// target从[0...target]因此，内层数组的容量应该为target+1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}
	// 递归边界
	dp[0][0] = 1
	// i从1开始
	// dp[i][c] = dp[i-1][c] + dp[i-1][c-nums[i-1]]
	for i := 1; i < len(dp); i++ {
		for c := 0; c < len(dp[0]); c++ {
			if c < nums[i-1] {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = dp[i-1][c] + dp[i-1][c-nums[i-1]]
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[len(nums)][target]

}

func TestFindTargetSumWays(t *testing.T) {
	//fmt.Println(findTargetSumWays1([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays2([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays3([]int{1, 1, 1, 1, 1}, 3))
}
