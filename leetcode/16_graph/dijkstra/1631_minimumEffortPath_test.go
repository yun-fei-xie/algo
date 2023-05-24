package dijkstra

import (
	"fmt"
	"sort"
	"testing"
)

/*
1631. 最小体力消耗路径
https://leetcode.cn/problems/path-with-minimum-effort/description/

你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row, col) 的高度。一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。
请你返回从左上角走到右下角的最小 体力消耗值 。

方法：将这m*n个节点放入并查集中，实时维护它们的连通性。
由于我们需要找到从左上角到右下角的最短路径，
因此我们可以将图中的所有边按照权值从小到大进行排序，并依次加入并查集中。
当我们加入一条权值为w的边之后，如果左上角和右下角从非连通状态变为连通状态，那么w即为答案。
*/
type edge struct {
	v1, v2, w int
}

func minimumEffortPath2(heights [][]int) int {
	var edges = make([]*edge, 0)
	// 建立图，heights数组中元素的个数就是图中节点的个数 这个题直接收集边
	row, col := len(heights), len(heights[0])
	// i是行号 j是列号
	for i, rowItems := range heights {
		for j, height := range rowItems {
			if i < row-1 {
				//生成向下的边 节点编号，坐标{i,j}-> i*col+j 坐标从0开始
				e := &edge{
					v1: i*col + j,
					v2: (i+1)*col + j,
					w:  abs(height, heights[i+1][j]),
				}
				edges = append(edges, e)
			}

			if j < col-1 {
				// 生成向右的边
				e := &edge{
					v1: i*col + j,
					v2: i*col + j + 1,
					w:  abs(height, heights[i][j+1]),
				}
				edges = append(edges, e)
			}
		}

	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	//
	uf := newUnionFind(col * row)
	for _, edge := range edges {
		uf.unionElements(edge.v1, edge.v2)
		if uf.isConnect(0, col*row-1) {
			return edge.w
		}
	}

	return -1
}

/*
并查集
*/
type unionFind struct {
	parent []int
	count  int
}

func newUnionFind(n int) *unionFind {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &unionFind{
		parent: p,
		count:  n,
	}
}

func (uf *unionFind) find(p int) int {
	return uf.parent[p]
}
func (uf *unionFind) isConnect(p, q int) bool {
	return uf.parent[p] == uf.parent[q]
}
func (uf *unionFind) unionElements(p int, q int) {
	pRoot := uf.parent[p]
	qRoot := uf.parent[q]
	for i := 0; i < uf.count; i++ {
		if uf.parent[i] == pRoot {
			uf.parent[i] = qRoot
		}
	}
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func TestMiniEffortPath(t *testing.T) {
	fmt.Println(minimumEffortPath2([][]int{{1, 10, 6, 7, 9, 10, 4, 9}}))
}
