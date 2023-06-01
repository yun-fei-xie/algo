package bridge

import (
	"fmt"
	"testing"
)

/*
1192. 查找集群内的关键连接
https://leetcode.cn/problems/critical-connections-in-a-network/?envType=study-plan-v2&envId=graph-theory

力扣数据中心有 n 台服务器，分别按从 0 到 n-1 的方式进行了编号。它们之间以 服务器到服务器 的形式相互连接组成了一个内部集群，连接是无向的。用  connections 表示集群网络，connections[i] = [a, b] 表示服务器 a 和 b 之间形成连接。任何服务器都可以直接或者间接地通过网络到达任何其他服务器。
关键连接 是在该集群中的重要连接，假如我们将它移除，便会导致某些服务器无法访问其他服务器。
请你以任意顺序返回该集群内的所有 关键连接 。

如果删除一条边，图中的连通分量发生了变化。那么这条边就是桥。
桥是两个强联通分量之间的连接边。
方法:Tarjan算法
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

func TestCriticalConnections(t *testing.T) {
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
	fmt.Println(criticalConnections(2, [][]int{{0, 1}}))
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 3}, {2, 3}}))
}
