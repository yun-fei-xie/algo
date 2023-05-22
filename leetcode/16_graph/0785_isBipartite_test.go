package _6_graph

import "container/list"

/*
785. 判断二分图
https://leetcode.cn/problems/is-graph-bipartite/description/

存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。给你一个二维数组 graph ，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图同时具有以下属性：
不存在自环（graph[u] 不包含 u）。
不存在平行边（graph[u] 不包含重复值）。
如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。
如果图是二分图，返回 true ；否则，返回 false 。

方法：通过深度优先遍历或者广度优先遍历对图进行染色
*/

/*
深度优先遍历
*/
func isBipartite1(graph [][]int) bool {
	var visited = make([]bool, len(graph))
	var color = make([]int, len(graph))

	// i：节点编号  c 颜色{0,1}
	var dfs func(i int, c int) bool
	dfs = func(i int, c int) bool {

		visited[i] = true
		color[i] = c

		for _, node := range graph[i] {
			// node已经被访问的情况
			if visited[node] && color[node] != (1-c) {
				return false
			}
			// node还没有被访问的情况
			if !visited[node] {
				if !dfs(node, 1-c) {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			if !dfs(i, 0) {
				return false
			}
		}
	}
	return true
}

/*
广度优先遍历
*/
func isBipartite2(graph [][]int) bool {
	var visited = make([]bool, len(graph))
	var color = make([]int, len(graph))

	// 放入节点放入队列之前，先访问并染色。从队列中取出节点是为了将其相邻节点进行访问和染色。
	// 所以，染色不能和访问不能放在出队的时候做。（如果放在出队做，第一个节点和后续的节点无法统一处理）
	var bfs func(i int, c int) bool
	bfs = func(i int, c int) bool {
		visited[i] = true
		color[i] = c

		queue := list.New()
		queue.PushBack(i)

		for queue.Len() != 0 {
			node := queue.Remove(queue.Front()).(int)
			for _, n := range graph[node] {
				if visited[n] == true && color[n] != 1-color[node] {
					return false
				}
				if !visited[n] {
					queue.PushBack(n)
					visited[n] = true
					color[n] = 1 - color[node]
				}
			}
		}
		return true
	}

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			if !bfs(i, 0) {
				return false
			}
		}
	}
	return true
}
