package treeDP

import (
	"fmt"
	"testing"
)

/*
2246. 相邻字符不同的最长路径
https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/description/
给你一棵 树（即一个连通、无向、无环图），根节点是节点 0 ，这棵树由编号从 0 到 n - 1 的 n 个节点组成。用下标从 0 开始、长度为 n 的数组 parent 来表示这棵树，其中 parent[i] 是节点 i 的父节点，由于节点 0 是根节点，所以 parent[0] == -1 。
另给你一个字符串 s ，长度也是 n ，其中 s[i] 表示分配给节点 i 的字符。
请你找出路径上任意一对相邻节点都没有分配到相同字符的 最长路径 ，并返回该路径的长度。
*/
func longestPath(parent []int, s string) int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		pa := parent[i]          // parent[i] 节点i的父亲节点 pa
		g[pa] = append(g[pa], i) // g[pa]表示表示节点pa的所有孩子节点
	}

	var ans int
	// 以节点i为子树root节点的相邻字符最长不同路径
	var dfs func(i int) int
	dfs = func(i int) int {
		// 遍历子树
		var maxLen int
		for _, child := range g[i] {

			len := dfs(child) + 1
			if s[i] != s[child] {
				ans = max(ans, maxLen+len)
				maxLen = max(maxLen, len)
			}

		}
		return maxLen
	}
	dfs(0)
	return ans
}

func TestLongestPath(t *testing.T) {

	// 一次遍历 求出最大+次大
	var arr = []int{7, 5, 3, 8, 6, 10}
	var m = 0
	var ans = 0
	for i := 0; i < len(arr); i++ {
		ans = max(ans, arr[i]+m)
		m = max(m, arr[i])
		fmt.Printf("ans->%d max->%d \n", ans, m)
	}
	fmt.Println(ans)
}
