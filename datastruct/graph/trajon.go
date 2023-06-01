package graph

import "fmt"

/*
使用深度优先遍历寻找图中的桥
leetcode 1192关键路径问题
*/
func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, connection := range connections {
		graph[connection[0]] = append(graph[connection[0]], connection[1])
		graph[connection[1]] = append(graph[connection[1]], connection[0])
	}

	visited := make([]bool, n)
	low := make([]int, n)
	order := make([]int, n)
	orderId := 0

	var ans = make([][]int, 0)

	var dfs func(v int, parent int)
	dfs = func(v int, parent int) {
		visited[v] = true
		order[v] = orderId
		orderId++
		low[v] = order[v]

		for _, adj := range graph[v] {

			if !visited[adj] {
				dfs(adj, v)
				// 判断v->adj是不是桥
				if low[adj] <= order[v] {
					//low[v] = low[adj]
					low[v] = min(low[v], low[adj])
				} else {
					ans = append(ans, []int{v, adj})
				}
			} else {
				if adj != parent {
					low[v] = min(low[adj], low[v])
				}
			}
		}

	}

	dfs(0, 0)
	fmt.Println(low)
	fmt.Println(order)
	return ans
}
func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if m > args[i] {
			m = args[i]
		}
	}
	return m
}
