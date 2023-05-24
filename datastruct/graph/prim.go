package graph

import "math"

/*
Prim最小生成树算法
假设图有n个顶点，并且是一个没有重边的连通图
1.每次在横切边中找一条最短边
2.然后更新两个阵营（用一个visited bool数组）
*/
func Prim(n int, edges []WeightEdge) (mst []WeightEdge) {
	//0.创建图
	graph := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		graph[i] = map[int]int{}
	}
	for _, edge := range edges {
		graph[edge.VertexI][edge.VertexJ] = edge.Weight
	}

	//1.创建布尔数组visited 并放入一个起始点。这里从0开始搜索。（起始点是哪个顶点并不重要）
	visited := make([]bool, n)
	visited[0] = true

	// 循环n-1次，把剩下的n-1个顶点收录进来
	for i := 1; i < n; i++ {

		// 寻找最短横切边的过程可以进行优化 暴力双重循环比较低效
		minEdge := NewWeightEdge(-1, -1, math.MaxInt)
		for v := 0; v < n; v++ {
			if visited[v] {
				for w, l := range graph[v] {
					if !visited[w] && l < minEdge.Weight {
						minEdge = NewWeightEdge(v, w, l)
					}
				}
			}
		}
		mst = append(mst, minEdge)
		visited[minEdge.VertexI] = true
		visited[minEdge.VertexJ] = true
	}
	return mst
}
