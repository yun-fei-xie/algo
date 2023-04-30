package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

func findNumberOfLIS(nums []int) int {

	dp := make([]int, len(nums)) // 存储以位置i开头的最长递增子序列的长度
	ll := make([]int, len(nums)) // 存储以位置i开头的最长递增子序列的个数
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		ll[i] = 1
	}
	maxL := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				//更新
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					ll[i] = ll[j]
				} else if dp[i] == dp[j]+1 {
					ll[i] += ll[j]
				}
			}
		}
		maxL = max(maxL, dp[i])
	}
	var ans int
	for i := 0; i < len(dp); i++ {
		if dp[i] == maxL {
			ans += ll[i]
		}
	}

	return ans
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TestFindNumberOfLIS(t *testing.T) {
	//fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	//fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	//fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 3, 2}))
	//dp[i]  2 1 1  ll[i]  1 1
}
