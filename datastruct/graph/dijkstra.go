package graph

import (
	"fmt"
	"math"
	"testing"
)

/*
返回起点到各个点的单源最短路径
顶点编号是[0...n] n+1个顶点

时间复杂度是n^2
*/

/*
dijkstra1函数的性能瓶颈在
*/
func dijkstra1(n int, edges [][]int, weight []int, start int) []int {
	// 构建带权图 使用map数组
	graph := make([]map[int]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make(map[int]int)
	}
	for i := 0; i < len(edges); i++ {
		out := edges[i][0]
		in := edges[i][1]
		w := weight[i]
		graph[out][in] = w
	}

	// dijkstra 算法
	dist := make([]int, n+1)
	visited := make([]bool, n+1)

	for i := 0; i <= n; i++ {
		dist[i] = math.MaxInt
	}
	dist[start] = 0

	for {
		// 找还没有确定最短距离的最小值
		d := math.MaxInt
		v := -1
		for i := 0; i <= n; i++ {
			if !visited[i] && dist[i] < d {
				d = dist[i]
				v = i
			}
		}
		// 如果找不到的话，表示所有的点都计算出了最短路径，此时退出循环
		if v == -1 {
			break
		}
		// 首先确定s->v的最短距离的长度就是dist[v]
		visited[v] = true
		// 然后通过v和相邻的点的边进行一轮更新
		for adj, w := range graph[v] {
			if !visited[adj] {
				if dist[v]+w < dist[adj] {
					dist[adj] = dist[v] + w
				}
			}
		}
	}
	return dist
}

/*
使用优先队列查找当前的最小dis值对应的顶点
*/
func dijkstra2(n int, edges [][]int, weight []int, start int) []int {
	//todo
	return nil
}

func TestDijkstra(t *testing.T) {
	fmt.Println(dijkstra1(2, [][]int{{0, 1}, {1, 2}, {0, 2}}, []int{1, 2, 3}, 0))
}
