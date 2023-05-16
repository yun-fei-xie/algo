package lcs

import (
	"fmt"
	"testing"
)

/*
712. 两个字符串的最小ASCII删除和
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

方法：
和最长公共子序列的思考方向是一致的。
如果word1[i]==word2[j] ，那么不需要做任何操作。

如果word1[i]!=word2[j],那么考虑一下三种编辑方式：
1.删除word1[i]，相当于去掉word1[i]， dfs(i,j) = dfs(i-1 ,j) + 1
2.在word1[i]中插入一个字符，插入的字符肯定是等于word2[j],dfs(i,j)=dfs(i,j-1)+1
3.置换word1[i]，将word1[i]直接改成word2[j]，dfs(i,j)= dfs(i-1 ,j-1) + 1

如果word1[i]==word2[j]的话，当前就不需要编辑：
dfs(i,j) = dfs(i-1 ,j-1)

考虑边界情况：
当i<0的时候，word1为空字符，word2还有j个字符，所以最小需要编辑j次。
当j<0的时候，word2为空字符，word1还有i个字符，所以最小需要编辑i次。

思考题：打印出最小编辑过程。

*/

/*
递归写法
*/

func minDistanceII1(word1 string, word2 string) int {

	var dfs func(i int, j int) int
	// [0...i]->长度i+1 [0...j]长度j+1
	dfs = func(i int, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		if word1[i] == word2[j] {
			return dfs(i-1, j-1)
		} else {
			//替换word1[i]，则替换后word1[i]==word2[j]，此时需要看word1[0...i-1]和word2[0...j-1]最小编辑距离
			//添加，在word1[i]后面添加一个元素，则添加后word1[i]==word2[j]，此时需要看word1[0...i]与word2[0...j-1]的最小编辑距离
			//删除，删除word1[i]。则此时需要知道word1[0...i-1]和word2[0...j]之间的最小编辑距离。
			return min(dfs(i-1, j-1), dfs(i, j-1), dfs(i-1, j)) + 1
		}
	}
	return dfs(len(word1)-1, len(word2)-1)
}

func minDistanceII2(word1 string, word2 string) int {
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

func minDistanceII3(word1 string, word2 string) int {
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

func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

func TestMinDistance72(t *testing.T) {
	fmt.Println(minDistanceII1("horse", "ros"))
	fmt.Println(minDistanceII2("horse", "ros"))
	fmt.Println(minDistanceII3("horse", "ros"))
}
