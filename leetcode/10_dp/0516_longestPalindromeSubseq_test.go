package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/longest-palindromic-subsequence/description/

给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：
输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
*/

/*
动态规划
dp[i][j]表示s[i...j]这段区间的最长会问子序列的长度
*/
func longestPalindromeSubseq(s string) int {
	lens := len(s)
	dp := make([][]int, lens)
	for i := 0; i < lens; i++ {
		dp[i] = make([]int, lens)
		dp[i][i] = 1 // 两个位置一样时，一定是回文
	}

	for i := lens - 1; i >= 0; i-- {
		for j := i + 1; j < lens; j++ {

			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[0][lens-1]
}

func TestLongestPalindromeSubSeq(t *testing.T) {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
