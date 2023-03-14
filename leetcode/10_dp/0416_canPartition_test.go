package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/partition-equal-subset-sum/

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

思考：
解法1：计算出数组的和sum（sum必须为偶数），然后用回溯找出有没有几个元素的和=sum/2  会超时

解法2：动态规划（和打家劫舍有点类似：偷还是不偷）
*/

func canPartition2(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	// 本质是一个0-1背包问题，包的容量为 sum/2 , 物品i的重量是 nums[i]，每一次检查i,i都可要可不要
	// 搜索是否有一个位置可以恰好装满整个背包
	dp := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		dp = append(dp, make([]int, sum/2+1)) // [0...sum/2]
	}
	// 背包容量为0
	for i := 0; i < len(nums); i++ {
		dp[i][0] = 0
	}
	// 只放第0号位置的元素
	for i := 1; i <= sum/2; i++ {
		if nums[0] > i {
			dp[0][i] = 0
		} else {
			dp[0][i] = nums[0]
		}
	}

	for j := 1; j < len(nums); j++ {

	}

	return false
}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}

	res := false

	var dfs func(arr []int, startIndex int, cap int)
	dfs = func(arr []int, startIndex int, cap int) {
		if startIndex >= len(arr) || cap < 0 {
			return
		}
		if cap == 0 { // startIndex <len(arr) && cap ==0
			res = true
			return
		}

		dfs(arr, startIndex+1, cap-arr[startIndex])
		dfs(arr, startIndex+1, cap)

	}
	dfs(nums, 0, sum/2)
	return res
}

func TestCanPartition(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	nums2 := []int{1, 2, 3, 5}
	res := canPartition(nums)
	res2 := canPartition(nums2)
	fmt.Println(res)
	fmt.Println(res2)

}
