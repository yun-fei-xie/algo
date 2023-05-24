package mst

import (
	"fmt"
	"sort"
	"testing"
)

/*
1584. 连接所有点的最小费用
https://leetcode.cn/problems/min-cost-to-connect-all-points/description/
给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。
连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。
请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。

方法：最小生成树算法 这一题我用kruskal算法 它用的是并查集
*/
type WeightEdge struct {
	VertexI int
	VertexJ int
	Weight  int
}

func minCostConnectPoints(points [][]int) int {
	// 1.二重循环生成边
	n := len(points)
	edges := make([]WeightEdge, 0)
	for i := 0; i < n; i++ {
		for j := 0; j != i && j < n; j++ {
			dist := abs(points[i][0], points[j][0]) + abs(points[i][1], points[j][1])
			edges = append(edges, WeightEdge{
				VertexI: i,
				VertexJ: j,
				Weight:  dist,
			})
		}
	}

	// 2.对边进行排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	//2.使用并查集
	uf := NewUnionFind(len(points))
	var ans int
	for i := 0; i < len(edges); i++ {
		if !uf.isConnect(edges[i].VertexI, edges[i].VertexJ) {
			ans += edges[i].Weight
			uf.unionElements(edges[i].VertexI, edges[i].VertexJ)
		}
	}
	return ans
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

/*
并查集部分
*/

type unionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *unionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &unionFind{
		parent: parent,
		count:  n,
	}
}

/*
查找节点p的父亲节点
*/
func (uf *unionFind) find(p int) int {
	return uf.parent[p]
}

/*
判断两个节点是否连通
*/
func (uf *unionFind) isConnect(p, q int) bool {
	return uf.parent[p] == uf.parent[q]
}

/*
合并两个节点
*/
func (uf *unionFind) unionElements(p, q int) {
	pId := uf.parent[p]
	qId := uf.parent[q]
	if pId == qId {
		return
	}
	for i := 0; i < uf.count; i++ {
		if uf.parent[i] == pId {
			uf.parent[i] = qId
		}
	}
}

func TestMinCostConnectPoints(t *testing.T) {
	//fmt.Println(minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}))
	//fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
	//fmt.Println(minCostConnectPoints([][]int{{0, 0}}))
	//fmt.Println(minCostConnectPoints([][]int{{-1000000, -1000000}, {1000000, 1000000}}))
	//fmt.Println(minCostConnectPoints([][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}))

	fmt.Println(minCostConnectPoints([][]int{{2, -3}, {-17, -8}, {13, 8}, {-17, -15}}))
}
