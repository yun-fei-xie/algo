package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/distinct-subsequences/description/?orderBy=hot
给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数。
输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit
*/

/*
把题目翻译过来，就相当于从s中按照从左到右的顺序挑选字符，组成的字符串等于字符串t的个数。
这样思考，这个问题就变得容易理解了。
每个字符都有两种状态，选择与不选择。
*/

func numDistinct(s string, t string) int { //s是主串、t是模式串
	lenS, lenT := len(s), len(t)

	var dfs func(i, j int) int //
	dfs = func(i, j int) int {
		if j < 0 {
			return 1
		}
		if i < 0 {
			return 0
		}
		if s[i] == t[j] {
			return dfs(i-1, j-1) + dfs(i-1, j)
		} else {
			return dfs(i-1, j)
		}
	}
	return dfs(lenS-1, lenT-1)
}

func numDistinct2(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	// 初始化
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}
	// dp[0][j] 为 0，默认值，因此不需要初始化
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
func TestNumDistinct(t *testing.T) {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	//fmt.Println(numDistinct2("rabbbit", "rabbit"))
}
