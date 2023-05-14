package _45

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/count-the-number-of-complete-components/
6432. 统计完全连通分量的数量

给你一个整数 n 。现有一个包含 n 个顶点的 无向 图，顶点按从 0 到 n - 1 编号。给你一个二维整数数组 edges 其中 edges[i] = [ai, bi] 表示顶点 ai 和 bi 之间存在一条 无向 边。
返回图中 完全连通分量 的数量。
如果在子图中任意两个顶点之间都存在路径，并且子图中没有任何一个顶点与子图外部的顶点共享边，则称其为 连通分量 。
如果连通分量中每对节点之间都存在一条边，则称其为 完全连通分量 。

方法：深度优先遍历。
在深度优先遍历的过程中用两个全局变量v和e记录一轮遍历下来，访问到的顶点和边的数量。
完全连通分量中的每一个都是一个完全的子图。满足每一个顶点和其他顶点之间都有边的关系。
n*(n-1)/2 = e
*/
func countCompleteComponents(n int, edges [][]int) int {
	var length = len(edges)
	var graph = make([][]int, n)
	for i := 0; i < length; i++ {
		x, y := edges[i][0], edges[i][1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	var visit = make([]bool, n)

	// 深度优先遍历
	var v, e int
	var dfs func(i int)
	dfs = func(i int) {
		visit[i] = true
		v++
		e += len(graph[i])

		for _, node := range graph[i] {
			if visit[node] == false {
				dfs(node)
			}
		}
	}
	// n(n-1)/2 = e
	var ans int
	for i := 0; i < n; i++ {
		v, e = 0, 0
		if visit[i] == false {
			dfs(i)
			if e == v*(v-1) {
				ans++
			}
		}
	}
	return ans
}

func TestCountComplete(t *testing.T) {
	//fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
	fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}))
}
