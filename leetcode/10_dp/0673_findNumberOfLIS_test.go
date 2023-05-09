package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
在第53号问题：最长上升子序列的基础上增加了要求，需要在求解的过程中记录下最长上升子序列的个数。
dp[i]记录以nums[i]为结尾的最长递增子序列的长度。 初始时都为1。
ll[i]记录以nums[i]为结尾的最长递增子序列的个数。 初始时都为1。

ll 如何更新？
if dp[i] < dp[j]+1 -> 说明出现了**新**的最长子序列，更新ll[i]。
此时，ll[i] = ll[j]。 -> 类比路径问题

else if dp[i] == dp[j]+1 -> 说明再次出现了一个最长子序列，累加ll[i]。
此时，ll[i] += ll[j]。-> 同样类别路径问题，出现了同样的最长子序列。

最后，遍历ll[i]收集结果。
*/
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
func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

func TestFindNumberOfLIS(t *testing.T) {
	//fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	//fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	//fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 3, 2}))
	//dp[i]  2 1 1  ll[i]  1 1
}
