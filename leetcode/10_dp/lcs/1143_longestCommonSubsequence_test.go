package lcs

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/longest-common-subsequence/
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。

这道题和718号问题
https://leetcode.cn/problems/maximum-length-of-repeated-subarray/
有点像。
更像是300、674、718三个题的结合版本。

dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]

单个数组或者字符串要用动态规划时，可以把动态规划 dp[i] 定义为 nums[0:i] 中想要求的结果；当两个数组或者字符串要用动态规划时，可以把动态规划定义成两维的 dp[i][j] ，其含义是在 A[0:i] 与 B[0:j] 之间匹配得到的想要的结果。

当 text1[i - 1] == text2[j - 1] 时，说明两个子字符串的最后一位相等，所以最长公共子序列又增加了 1，所以 dp[i][j] = dp[i - 1][j - 1] + 1；举个例子，比如对于 ac 和 bc 而言，他们的最长公共子序列的长度等于 a和 b的最长公共子序列长度 0 + 1 = 1。
当 text1[i - 1] != text2[j - 1] 时，说明两个子字符串的最后一位不相等，那么此时的状态 dp[i][j] 应该是 dp[i - 1][j] 和 dp[i][j - 1] 的最大值。举个例子，比如对于 ace和bc而言，他们的最长公共子序列的长度等于 ① ace 和 b 的最长公共子序列长度0 与 ② ac 和 bc的最长公共子序列长度1 的最大值，即 1。

	a  b  c  d  e

a
c
e
*/

/*
递归写法
*/
func longestCommonSubSequence1(text1 string, text2 string) int {

	m := len(text1)
	n := len(text2)

	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i < 0 || j < 0 {
			return 0
		}

		if text1[i] == text2[j] {
			return dfs(i-1, j-1) + 1
		}
		return max(dfs(i-1, j), dfs(i, j-1))
	}
	return dfs(m, n)
}

/*
1:1翻译成递推
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len2+1)
	for i := 0; i <= len2; i++ {
		dp[i] = make([]int, len1+1)
	}

	for i := 1; i <= len2; i++ {
		for j := 1; j <= len1; j++ {

			if text2[i-1] == text1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// 最大值肯定是右下角 因为在求解的过程中，不是+1，就是取max
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[len2][len1]
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

func TestLongestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	//fmt.Println(longestCommonSubsequence("abc", "abc"))
	//fmt.Println(longestCommonSubsequence("ac", "abc"))
}
