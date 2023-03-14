package mid

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/word-break/description/?favorite=2cktkvj

给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

回溯算法+记忆化  题目会出现重叠子问题
DFS 思路
"leetcode"能否 break，可以拆分为：
"l"是否是单词表的单词、剩余子串能否 break。
"le"是否是单词表的单词、剩余子串能否 break。
"lee"...以此类推
用 DFS 回溯，考察所有的拆分可能，指针从左往右扫描：
如果指针的左侧部分是单词，则对剩余子串递归考察。
如果指针的左侧部分不是单词，不用看了，回溯，考察别的分支。
*/
func wordBreak(s string, wordDict []string) bool {

	mp := make(map[string]struct{})
	for i := 0; i < len(wordDict); i++ {
		mp[wordDict[i]] = struct{}{}
	}

	memo := make(map[int]bool, 0)

	var dfs func(str string, startIndex int) bool

	// 考察从[startIndex , len(str)-1]是否满足条件
	dfs = func(str string, startIndex int) bool {
		// 考察[startIndex , len(str)-1]
		if startIndex >= len(str) { // 空值返回true
			return true
		}

		if b, found := memo[startIndex]; found {
			return b
		}

		for i := startIndex; i < len(str); i++ {
			_, found := mp[str[startIndex:i+1]]
			if found {
				if b, found := memo[i+1]; found {
					if b == true {
						return true
					}
				} else {
					b := dfs(str, i+1)
					memo[i+1] = b
					if b == true {
						return true
					}
				}
			}
		}
		memo[startIndex] = false
		return memo[startIndex]
	}

	return dfs(s, 0)
}

func TestWordBreak(t *testing.T) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	res := wordBreak(s, wordDict)
	fmt.Println(res)
}
