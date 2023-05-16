package lcs

import (
	"fmt"
	"testing"
)

/*
712. 两个字符串的最小ASCII删除和
https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/
给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。



示例 1:

输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
示例 2:

输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。


提示:

0 <= s1.length, s2.length <= 1000
s1 和 s2 由小写英文字母组成

方法：
按照最长公共子序列的思考方式：
如果s1[i]==s2[j]，此时，不需要删除。最小ASCII dfs(i,j) = dfs(i-1 ,j-1)
如果s1[i]!=s2[j]，此时
	1. 考虑删除s1[i]，此时最小ASCII dfs(i,j) = dfs(i-1 ,j) + ASCII(s1[i])
	2. 考虑删除s2[j]，此时最小ASCII dfs(i,j) = dfs(i ,j-1) + ASCII(s2[j])
    3. 两个都删除，此时最小ASCII dfs(i,j) = dfs(i-1 ,j-1) + ASCII(s1[i]) + ASCII(s2[j])
	取三者的最小值
*/

/*
递归写法
*/
func minimumDeleteSum1(s1 string, s2 string) int {

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return sumASCII(0, j, &s2)
		}
		if j < 0 {
			return sumASCII(0, i, &s1)
		}

		if s1[i] == s2[j] {
			return dfs(i-1, j-1)
		} else {
			return min(dfs(i-1, j)+int(s1[i]), dfs(i, j-1)+int(s2[j]), dfs(i-1, j-1)+int(s1[i])+int(s2[j]))
		}
	}
	return dfs(len(s1)-1, len(s2)-1)

}

func sumASCII(left, right int, s *string) (sum int) {

	for left <= right {
		sum += int((*s)[left])
		left++
	}
	return sum
}

/*
1:1翻译成递归
*/

func minimumDeleteSum2(s1 string, s2 string) int {

	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = sumASCII(0, j-1, &s2)
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = sumASCII(0, i-1, &s1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]), dp[i-1][j-1]+int(s1[i-1])+int(s2[j-1]))
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func TestMinimumDeleteSum(t *testing.T) {
	fmt.Println(minimumDeleteSum1("sea", "eat"))
	fmt.Println(minimumDeleteSum2("sea", "eat"))
}
