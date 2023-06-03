package unionFind_test

/*
1579. 保证图可完全遍历
https://leetcode.cn/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/description/

Alice 和 Bob 共有一个无向图，其中包含 n 个节点和 3  种类型的边：
类型 1：只能由 Alice 遍历。
类型 2：只能由 Bob 遍历。
类型 3：Alice 和 Bob 都可以遍历。
给你一个数组 edges ，其中 edges[i] = [typei, ui, vi] 表示节点 ui 和 vi 之间存在类型为 typei 的双向边。请你在保证图仍能够被 Alice和 Bob 完全遍历的前提下，找出可以删除的最大边数。如果从任何节点开始，Alice 和 Bob 都可以到达所有其他节点，则认为图是可以完全遍历的。
返回可以删除的最大边数，如果 Alice 和 Bob 无法完全遍历图，则返回 -1 。


想法：判断一条边是否能够被删除。如果删除这条边，联通分量发生了变化，那么这条边就不能被删除。
	有没有办法，快速判断尝试拿掉一条边之后，整个图的联通分量的变化情况。

方法：
1.先判断通用边,然后用两个并查集分别判断alice的边和bob的边

*/

func maxNumEdgesToRemove(n int, edges [][]int) int {
	uf1 := newNumEdgeUF(n)
	uf2 := newNumEdgeUF(n)
	length := len(edges)
	ans := 0
	for i := 0; i < length; i++ {
		if edges[i][0] == 3 {
			success := uf1.unionElements(edges[i][1]-1, edges[i][2]-1)
			uf2.unionElements(edges[i][1]-1, edges[i][2]-1)
			if !success {
				ans++
			}
		}
	}

	for i := 0; i < length; i++ {
		if edges[i][0] == 1 {
			if !uf1.unionElements(edges[i][1]-1, edges[i][2]-1) {
				ans++
			}
		}
		if edges[i][0] == 2 {
			if !uf2.unionElements(edges[i][1]-1, edges[i][2]-1) {
				ans++
			}
		}
	}

	if uf1.sizeCount != 1 || uf2.sizeCount != 1 {
		return -1
	}

	return ans
}

type numEdgeUF struct {
	parent    []int
	sizeCount int
}

func newNumEdgeUF(n int) *numEdgeUF {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &numEdgeUF{
		parent:    p,
		sizeCount: n,
	}
}

func (uf *numEdgeUF) find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *numEdgeUF) unionElements(p int, q int) bool {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return false
	}

	uf.parent[pRoot] = qRoot
	uf.sizeCount--
	return true
}
