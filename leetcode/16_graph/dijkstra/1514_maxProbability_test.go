package dijkstra

import (
	"fmt"
	"math"
	"testing"
)

/*
1514. 概率最大的路径
https://leetcode.cn/problems/path-with-maximum-probability/description/
给你一个由 n 个节点（下标从 0 开始）组成的无向加权图，该图由一个描述边的列表组成，其中 edges[i] = [a, b] 表示连接节点 a 和 b 的一条无向边，且该边遍历成功的概率为 succProb[i] 。
指定两个节点分别作为起点 start 和终点 end ，请你找出从起点到终点成功概率最大的路径，并返回其成功概率。
如果不存在从 start 到 end 的路径，请 返回 0 。只要答案与标准答案的误差不超过 1e-5 ，就会被视作正确答案。

方法：迪杰斯特拉算法
*/
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	//1.建图
	graph := make([]map[int]float64, n)
	for i := 0; i < n; i++ {
		graph[i] = map[int]float64{}
	}
	for index, edge := range edges {
		out := edge[0]
		in := edge[1]
		graph[out][in] = succProb[index]
		graph[in][out] = succProb[index]
	}

	//2.visited数组将图分为两个部分
	visited := make([]bool, n)

	//3.dist数组 题目要求找最大值，所以初始值都赋值为最小
	dist := make([]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MinInt
	}
	// 起点到起点的概率应该是1
	dist[start] = 1

	for {

		// 找还没有确定最长距离的最大值
		d := float64(math.MinInt)
		v := -1
		for i := 0; i < n; i++ {
			if !visited[i] && dist[i] > d {
				d = dist[i]
				v = i
			}
		}
		if v == -1 {
			break
		}

		visited[v] = true
		// 根据v更新相邻的顶点
		for adj, w := range graph[v] {
			if dist[adj] < dist[v]*w {
				dist[adj] = dist[v] * w
			}
		}
	}
	if dist[end] == math.MinInt {
		return 0.00000
	}
	return dist[end]
}

func TestMaxProbability(t *testing.T) {
	fmt.Println(maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
	fmt.Println(maxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2))
}
