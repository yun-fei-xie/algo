package main

import (
	"fmt"
	"sort"
)

/*
建立一个虚拟的节点0,
需要自建的表示从节点0过去
*/
func mainp4() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		// n家企业、m家合作关系 扩从到n+1家（添加虚拟vertex 0）
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		// 使用kruskal只需要边的集合
		edges := make([]*WeightEdge, 0)
		// 虚拟节点0与每家企业之间的边
		var w int
		for k := 1; k <= n; k++ {
			fmt.Scanf("%d", &w)
			edges = append(edges, &WeightEdge{VertexI: 0, VertexJ: k, Weight: w})
			edges = append(edges, &WeightEdge{VertexI: k, VertexJ: 0, Weight: w})
		}
		// 真实企业之间的边
		var e1, e2, cost int
		for k := 0; k < m; k++ {
			fmt.Scanf("%d %d %d", &e1, &e2, &cost)
			edges = append(edges, &WeightEdge{VertexI: e1, VertexJ: e2, Weight: cost})
			edges = append(edges, &WeightEdge{VertexI: e2, VertexJ: e1, Weight: cost})
		}
		// kruskal算法求解最小生成树的边的权重和
		fmt.Println(minCost(edges, n+1))

	}
}

type WeightEdge struct {
	VertexI int
	VertexJ int
	Weight  int
}

// kruskal算法求解最少生成树
func minCost(edges []*WeightEdge, cnts int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	var ans int
	uf := InitUnionFind(cnts)
	for _, edge := range edges {
		if !uf.IsConnected(edge.VertexI, edge.VertexJ) {
			ans += edge.Weight
			uf.UnionElements(edge.VertexI, edge.VertexJ)
		}
	}
	return ans
}

// 并查集

type UnionFind struct {
	id    []int
	count int
}

// InitUnionFind 初始时，每个元素各自为一组，互不连接
func InitUnionFind(n int) *UnionFind {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &UnionFind{
		id:    id,
		count: n,
	}
}

func (un *UnionFind) Find(p int) int {
	return un.id[p]
}

func (un *UnionFind) IsConnected(p, q int) bool {
	return un.id[p] == un.id[q]
}

func (un *UnionFind) UnionElements(p, q int) {
	pId := un.Find(p)
	qId := un.Find(q)
	if pId == qId {
		return
	}
	for i := 0; i < un.count; i++ {
		if un.id[i] == pId {
			un.id[i] = qId
		}
	}
}
