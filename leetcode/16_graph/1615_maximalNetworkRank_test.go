package _6_graph_test

/*
1615. 最大网络秩
https://leetcode.cn/problems/maximal-network-rank/description/?envType=study-plan-v2&envId=graph-theory

n 座城市和一些连接这些城市的道路 roads 共同组成一个基础设施网络。每个 roads[i] = [ai, bi] 都表示在城市 ai 和 bi 之间有一条双向道路。
两座不同城市构成的 城市对 的 网络秩 定义为：与这两座城市 直接 相连的道路总数。如果存在一条道路直接连接这两座城市，则这条道路只计算 一次 。
整个基础设施网络的 最大网络秩 是所有不同城市对中的 最大网络秩 。
给你整数 n 和数组 roads，返回整个基础设施网络的 最大网络秩 。

方法：邻接表建图，统计顶点i和顶点j的出边总和rank。如果<i,j>之间有一条通路的话，那么rank--（重复统计）
*/
func maximalNetworkRank(n int, roads [][]int) int {

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], road[1])
		graph[road[1]] = append(graph[road[1]], road[0])
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			var hasRoad bool
			for _, adj := range graph[i] {
				if adj == j {
					hasRoad = true
					break
				}
			}
			rank := len(graph[i]) + len(graph[j])
			if hasRoad {
				rank--
			}
			if rank > ans {
				ans = rank
			}
		}
	}
	return ans
}
