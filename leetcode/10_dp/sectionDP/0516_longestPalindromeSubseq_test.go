package sectionDP

import (
	"fmt"
	"testing"
)

/*
516.最长回文子序列
https://leetcode.cn/problems/longest-palindromic-subsequence/description/

给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：
输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。

方法：将s反转为s1，然后求解s与s1的最长回文子序列

方法：递归，从左右两侧向中间考虑。每次考虑两侧两个字符（选择还是不选）这个和最长公共子序列的思考方式本质是一样的。

	有三种情况：如果s[i]==s[j] dfs(i,j)= 2 + dfs(i+1 ，j-1) 两个都选
			如果s[i]!=s[j] , 再分为三种情况，选左边不选右边，选择右边不选择左边 左右都不选 （左右都不选可以去掉，它包含在前面两种情况种）
	baseCase：如果i>j 返回0(没有字母) ，如果i==j 返回1(只有一个字母)

方法：区间dp
dp[i][j]表示s[i...j]这段区间的最长回文子序列的长度
*/

/*
递归代码
*/
func longestPalindromeSubseq1(s string) int {

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}
		if s[i] == s[j] {
			return dfs(i+1, j-1) + 2
		} else {
			return max(dfs(i+1, j), dfs(i, j-1), dfs(i+1, j-1))
		}
	}
	return dfs(0, len(s)-1)
}

/*
将递归代码1：1翻译成递推
*/
func longestPalindromeSubseq2(s string) int {
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

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

func TestLongestPalindromeSubSeq(t *testing.T) {
	fmt.Println(longestPalindromeSubseq1("bbbab"))
	fmt.Println(longestPalindromeSubseq2("bbbab"))
}
