package _0_dp

import (
	"fmt"
	"testing"
)

/*
判断子序列
https://leetcode.cn/problems/is-subsequence/description/?orderBy=hot

输入：s = "abc", t = "ahbgdc"
输出：true
示例 2：

输入：s = "axc", t = "ahbgdc"
输出：false

动态规划
*/
func isSubsequence(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	dp := make([][]int, lenS+1)
	for i := 0; i <= lenS; i++ {
		dp[i] = make([]int, lenT+1)
	}

	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenT; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[lenS][lenT] == len(s)

}

/*
双指针
指针1指向s
指针2指向t
*/
func isSubsequence2(s string, t string) bool {
	lens, lent := len(s), len(t)
	p1, p2 := 0, 0
	for p1 < lens && p2 < lent {
		if s[p1] == t[p2] {
			p1++
			p2++
		} else {
			p2++
		}
	}
	return p1 == lens
}

func TestIsSubSequence(t *testing.T) {
	fmt.Println(isSubsequence("abc", "acbcde"))
}
