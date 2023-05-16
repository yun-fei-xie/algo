package lcs

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/delete-operation-for-two-strings/

给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
每步 可以删除任意一个字符串中的一个字符。

输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"




*/

/*
解法1：转换为最长公共子序列问题。
通过观察可以发现，res = len(word1)+ len(word2) - 2 * (commonSubArray)
*/
func minDistance(word1 string, word2 string) int {
	lenWord1, lenWord2 := len(word1), len(word2)
	dp := make([][]int, lenWord1+1)
	for i := 0; i <= lenWord1; i++ {
		dp[i] = make([]int, lenWord2+1)
	}
	var maxCommLen int
	for i := 1; i <= lenWord1; i++ {
		for j := 1; j <= lenWord2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			maxCommLen = max(maxCommLen, dp[i][j])
		}
	}
	return lenWord1 + lenWord2 - 2*maxCommLen
}

/*
dp[i][j]：以i-1为结尾的字符串word1，和以j-1位结尾的字符串word2，想要达到相等，所需要删除元素的最少次数。
*/
func minDistance2(word1 string, word2 string) int {
	lenWord1, lenWord2 := len(word1), len(word2)
	dp := make([][]int, lenWord1+1)
	for i := 0; i <= lenWord1; i++ {
		dp[i] = make([]int, lenWord2+1)
	}
	// 初始化第一行，牢记需要删除的数量
	for i := 1; i <= lenWord2; i++ {
		dp[0][i] = i
	}
	// 初始化第一列，牢记需要删除的数量
	for j := 1; j <= lenWord1; j++ {
		dp[j][0] = j
	}

	for i := 1; i <= lenWord1; i++ {
		for j := 1; j <= lenWord2; j++ {
			// 如果两个字符相同，则不需要删除任何一个字符
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 如果两个字符不同，则需要删除至少一个字符
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[len(word1)][len(word2)]
}

func TestMinDistance(t *testing.T) {
	//fmt.Println(minDistance("sea", "eat"))
	//fmt.Println(minDistance("leetcode", "etco"))
	//fmt.Println(minDistance("park", "spake"))
	fmt.Println(minDistance2("park", "spake"))

}
