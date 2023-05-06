package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/edit-distance/

给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

示例 1：
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

*/

func minDistance72(word1 string, word2 string) int {
	lenWord1, lenWord2 := len(word1), len(word2)
	dp := make([][]int, lenWord1+1)
	for i := 0; i <= lenWord1; i++ {
		dp[i] = make([]int, lenWord2+1)
	}
	// 初始化行

	for i := 1; i <= lenWord2; i++ {
		dp[0][i] = i
	}
	// 初始化列
	for j := 1; j <= lenWord1; j++ {
		dp[j][0] = j
	}

	for i := 1; i <= lenWord1; i++ {
		for j := 1; j <= lenWord2; j++ {
			// 不需要编辑
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 需要编辑
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	//for i := 0; i < len(dp); i++ {
	//	fmt.Println(dp[i])
	//}
	return dp[lenWord1][lenWord2]
}

func minDistance3(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i // word1[i] 变成 word2[0], 删掉 word1[i], 需要 i 部操作
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j // word1[0] 变成 word2[j], 插入 word1[j]，需要 j 部操作
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else { // Min(插入，删除，替换)
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}

	return dp[m][n]
}
func TestMinDistance72(t *testing.T) {
	fmt.Println(minDistance72("horse", "ros"))
	//fmt.Println(minDistance3("horse", "ros"))
}
