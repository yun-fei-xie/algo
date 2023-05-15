package _6_graph

import (
	"fmt"
	"testing"
)

/*
797. 所有可能的路径
https://leetcode.cn/problems/all-paths-from-source-to-target/

给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。


方法：和二叉树的所有路径基本是一致的

*/

func allPathsSourceTarget(graph [][]int) [][]int {
	var ans = make([][]int, 0)
	var path = make([]int, 0)
	path = append(path, 0)

	var n = len(graph)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n-1 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for j := 0; j < len(graph[i]); j++ {
			path = append(path, graph[i][j])
			dfs(graph[i][j])
			path = path[0 : len(path)-1]
		}

	}

	dfs(0)
	return ans
}

func TestAllPathSourceTarget(t *testing.T) {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
