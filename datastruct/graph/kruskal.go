package graph

import (
	"algo/datastruct/disjointSet"
	"sort"
)

/*
最小生成树算法

kruskal算法：
假设图有n个顶点

1,前提：图需要时联通的
2.将图中的所有的边从小到大进行排序，使用贪心的思想，每次选择当前最短的一条横切边。直到选满了n-1条。
3.当前选择的边，不能让已经选择的边的构成环。
4.解决问题3，需要用一个bool数组。如果当前选择的边的两个端点之前没有全部被选择过，那么就不会和之前选择的边中包含的顶点构成环。
*/

/*
传入一副图
params: n表示图的顶点的个数
params: weightEdges表示图的所有的带权边
使用并查集来判断新加入的边是否会给当前最小生成树中的图带来环
*/
func kruskal2(n int, edges []WeightEdge) (mst []WeightEdge) {
	// 1.判断图的连通性，如果不连通，则无法获得最小生成树，返回空值。这里为了简化，不做判断。
	// 2.将所有的带权边按照权重从小到大进行排序。
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	//初始化并查集
	uf := disjointSet.InitUnionFind(n)

	//收集结果
	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		if !uf.IsConnected(edge.VertexI, edge.VertexJ) {
			mst = append(mst, edge)
			uf.UnionElements(edge.VertexI, edge.VertexJ)
		}
	}
	return mst
}
