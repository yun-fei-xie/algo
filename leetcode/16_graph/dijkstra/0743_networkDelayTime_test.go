package dijkstra

import "math"

/*
743. 网络延迟时间
https://leetcode.cn/problems/network-delay-time/description/

有 n 个网络节点，标记为 1 到 n。
给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。
现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。

方法：标准的dijkstra算法，找所有最短路径中的最大值
如果最后有顶点没有被更新，还是初始的无穷大值，那么这个图就不是连通图
*/
func networkDelayTime(times [][]int, n int, k int) int {
	//1.建图
	graph := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		graph[i] = map[int]int{}
	}
	for _, time := range times {
		out := time[0] - 1
		in := time[1] - 1
		t := time[2]
		graph[out][in] = t
	}

	// 2.准备visited数组和dist数组
	visited := make([]bool, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}
	dist[k-1] = 0

	//3.执行松弛操作
	for {
		d := math.MaxInt
		v := -1
		for i := 0; i < n; i++ {
			if !visited[i] && dist[i] < d {
				d = dist[i]
				v = i
			}
		}
		if v == -1 {
			break
		}
		visited[v] = true
		for e, t := range graph[v] {
			if !visited[e] && dist[v]+t < dist[e] {
				dist[e] = dist[v] + t
			}
		}
	}
	var ans = math.MinInt
	for i := 0; i < len(dist); i++ {
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
