package lis

import (
	"fmt"
	"testing"
)

/*
1027. 最长等差数列
https://leetcode.cn/problems/longest-arithmetic-subsequence/description/

给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。
回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik] ，且 0 <= i1 < i2 < ... < ik <= nums.length - 1。并且如果 seq[i+1] - seq[i]( 0 <= i < seq.length - 1) 的值都相同，那么序列 seq 是等差的。

示例 1：
输入：nums = [3,6,9,12]
输出：4
解释：
整个数组是公差为 3 的等差数列。
示例 2：

输入：nums = [9,4,7,2,10]
输出：3
解释：
最长的等差子序列是 [4,7,10]。
示例 3：

输入：nums = [20,1,15,3,10,5,8]
输出：4
解释：
最长的等差子序列是 [20,15,10,5]。


提示：

2 <= nums.length <= 1000
0 <= nums[i] <= 500

方法：枚举。枚举的思路和之前做最长上升子序列比较像。300号问题。
dp[i][j]表示的是，以nums[i]为结尾，公差为j的最长等差子序列的长度。
对于公差j的范围来说，由于nums[i]在0到500之间，所以，公差的范围[-500,500] 长度为1001。
在构造数组的时候，j需要开一个长度1001数组[0,1000]，j的偏移量为500

*/

func longestArithSeqLength2(nums []int) int {
	// 以nums[i]为结尾的 公差为j的最长等差子序列的长度
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 1001)
	}
	var ans int
	for i := 0; i < len(dp); i++ {
		// 遍历公差从
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j] + 500 // 偏移一下
			dp[i][d] = dp[j][d] + 1
			ans = max(ans, dp[i][d])
		}
	}
	return ans + 1
}

func TestLongestArithSeqLength(t *testing.T) {
	//fmt.Println(longestArithSeqLength1([]int{3, 6, 9, 12}))
	fmt.Println(longestArithSeqLength2([]int{3, 6, 9, 12}))

	//fmt.Println(longestArithSeqLength1([]int{22, 8, 57, 41, 36, 46, 42, 28, 42, 14, 9, 43, 27, 51, 0, 0, 38, 50, 31, 60, 29, 31, 20, 23, 37, 53, 27, 1, 47, 42, 28, 31, 10, 35, 39, 12, 15, 6, 35, 31, 45, 21, 30, 19, 5, 5, 4, 18, 38, 51, 10, 7, 20, 38, 28, 53, 15, 55, 60, 56, 43, 48, 34, 53, 54, 55, 14, 9, 56, 52}))
	//fmt.Println(longestArithSeqLength2([]int{22, 8, 57, 41, 36, 46, 42, 28, 42, 14, 9, 43, 27, 51, 0, 0, 38, 50, 31, 60, 29, 31, 20, 23, 37, 53, 27, 1, 47, 42, 28, 31, 10, 35, 39, 12, 15, 6, 35, 31, 45, 21, 30, 19, 5, 5, 4, 18, 38, 51, 10, 7, 20, 38, 28, 53, 15, 55, 60, 56, 43, 48, 34, 53, 54, 55, 14, 9, 56, 52}))
	//fmt.Println(longestArithSeqLength1([]int{9, 4, 7, 2, 10}))
	//fmt.Println(longestArithSeqLength1([]int{20, 1, 15, 3, 10, 5, 8}))
}
