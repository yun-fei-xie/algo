package subset

import (
	"fmt"
	"testing"
)

/*
131. 分割回文串
https://leetcode.cn/problems/palindrome-partitioning/description/

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]

方法：子集问题，枚举逗号的位置
*/
func partition(s string) [][]string {

	var res = make([][]string, 0)
	var path = make([]string, 0)

	var dfs func(s string, startIndex int)
	dfs = func(s string, startIndex int) {
		// 如果startIndex 超出范围，表示已经找到一份可行的方案
		if startIndex >= len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := startIndex; i < len(s); i++ {
			if ispalindromic(s[startIndex : i+1]) { // slice[a,b] -> [a,b)  // 这里的判断还有有剪枝的作用
				path = append(path, s[startIndex:i+1]) // 把这一段子串放到path中，注意slice的范围是左闭右开
				dfs(s, i+1)
				path = path[0 : len(path)-1]
			}
		}
	}
	dfs(s, 0)
	return res

}

// 判断回文串
func ispalindromic(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func TestPartition(t *testing.T) {
	s := "aab"
	fmt.Println(partition(s))
}
