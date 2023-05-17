package sectionDP

import (
	"fmt"
	"testing"
)

/*
1312. 让字符串成为回文串的最少插入次数
https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/description/

给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
请你返回让 s 成为回文串的 最少操作次数 。
「回文串」是正读和反读都相同的字符串。

方法：区间dp
做过最长公共子序列和最长回文子序列之后，这个题非常容易。
从字符串的两端考虑
*/
func minInsertions1(s string) int {

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}
		if s[i] == s[j] {
			return dfs(i+1, j-1)
		} else {
			return min(dfs(i, j-1), dfs(i+1, j)) + 1
		}
	}
	return dfs(0, len(s)-1)
}

/*
1:1翻译成递推
需要小心的是数组下标越界的问题。
下面的写法可能会越界的下标 i+1 ,j-1
也就是当i==len(s)-1 会发生上溢 ，或者 j==0的时候 j-1会发生下溢出。
如果给dp数组左边多一列，下面多一行，就不会出现溢出。
*/
func minInsertions2(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}

	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			//[i...j]
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i+1][j]) + 1
			}
		}
	}
	return dp[0][length-1]
}

/*
手动处理溢出
*/
func minInsertions3(s string) int {
	length := len(s)
	dp := make([][]int, length+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, length)
	}
	// i对应数组中的i+1 j
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			//[i...j]
			if s[i] == s[j] {
				if j-1 < 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				if j-1 < 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i][j-1], dp[i+1][j]) + 1
				}
			}
		}
	}
	return dp[0][length-1]
}

func TestMinInsertions(t *testing.T) {
	fmt.Println(minInsertions1("mbadm"))
	fmt.Println(minInsertions3("leetcode"))
}
