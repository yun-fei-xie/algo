package _0_dp

import (
	"fmt"
	"testing"
)

/*
这道题的代码和1143，最大公共子序列的代码一模一样。
*/
func maxUncrossedLines(nums1 []int, nums2 []int) int {

	len1, len2 := len(nums1), len(nums2)
	dp := make([][]int, len2+1)
	for i := 0; i <= len2; i++ {
		dp[i] = make([]int, len1+1)
	}

	for i := 1; i <= len2; i++ {
		for j := 1; j <= len1; j++ {

			if nums2[i-1] == nums1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len2][len1]
}

func TestMaxUncrossedLines(t *testing.T) {
	fmt.Println(maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
	fmt.Println(maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
}
