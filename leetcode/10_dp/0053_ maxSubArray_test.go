package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/

/*
双重循环搜索i到i这个区间 sum [i...j] 的最大值
应该是最容易理解
*/
func maxSubArray(nums []int) int {
	var ans = math.MinInt32
	var length = len(nums)
	var sum func(nums []int, i int, j int) int
	sum = func(nums []int, i int, j int) int {
		var res int
		for ; i <= j; i++ {
			res += nums[i]
		}
		return res
	}

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			ans = max(ans, sum(nums, i, j))
		}
	}
	return ans
}

/*
动态规划

nums[-2,1,-3,4,-1,2,1,-5,4]
dp  [-2,1,-2,4, 3,5,6,1 ,5]

用贡献来理解这个问题。
如果一个数子append到连续子序列，得到的和反而不如自己本身
dp[i]表示以nums[i]为结尾的最大连续子数组的和 i一定是dp[i]区间的结尾。
i也可能是dp[j]的开头

这个解法也被称为Kadane算法
kadane算法的要求是，dp[i]一定是以nums[i]为结尾。
也就是说nums[i]这个元素一定需要被选取。
https://zh.wikipedia.org/wiki/%E6%9C%80%E5%A4%A7%E5%AD%90%E6%95%B0%E5%88%97%E9%97%AE%E9%A2%98#:~:text=Kadane%E7%AE%97%E6%B3%95%E6%89%AB%E6%8F%8F%E4%B8%80%E6%AC%A1%E6%95%B4%E4%B8%AA,%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%9A%84%E4%B8%80%E4%B8%AA%E4%BE%8B%E5%AD%90%E3%80%82
*/
func maxSubArray2(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	var ans = dp[0]
	//	fmt.Printf("dp[%d] -> %d\n", 0, dp[0])
	for i := 1; i < length; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		//	fmt.Printf("dp[%d] -> %d\n", i, dp[i])
		ans = max(ans, dp[i])
	}
	return ans
}

// 求解最小子序列和 怎么求
func minSubArray(nums []int) int {

	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	minSub := nums[0]
	for i := 1; i < length; i++ {
		dp[i] = min(dp[i-1]+nums[i], nums[i]) // dp[i]一定要以nums[i]为结尾
		minSub = min(minSub, dp[i])
	}
	return minSub
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(minSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
