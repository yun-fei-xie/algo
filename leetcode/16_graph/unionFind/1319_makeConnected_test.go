package unionFind_test

import (
	"fmt"
	"testing"
)

/*
1319. 连通网络的操作次数
https://leetcode.cn/problems/number-of-operations-to-make-network-connected/

方法：转换为求解图中的连通分量的个数count。最后的结果就是count-1。
1.并查集求解图的连通分量个数需要做一些改造。加入一个setCount变量。
2.初始时，setCount=n也就是节点的个数。
3.每当发生元素合并操作，并查集中的组数就会--。

通过给并查集中增加一个setCount属性就可以追踪并查集中连通分量的个数。
通过增加一个size数组可以追踪每个连通分量包含节点的个数。（虽然这个题用不上，但是也可以实现一下）
*/
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	uf := NewMyUf(n)
	for _, conn := range connections {
		uf.unionElements(conn[0], conn[1])
	}

	return uf.setCount - 1
}

type MyUnionFind struct {
	parent   []int
	size     []int
	setCount int
}

func NewMyUf(n int) *MyUnionFind {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &MyUnionFind{
		parent:   p,
		size:     s,
		setCount: n,
	}
}

func (uf *MyUnionFind) isConnect(p int, q int) bool {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	return pRoot == qRoot
}

// 合并的时候，将节点少的树挂载节点多的树上
func (uf *MyUnionFind) unionElements(p int, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	pSize := uf.size[pRoot]
	qSize := uf.size[qRoot]
	if pSize < qSize {
		uf.parent[pRoot] = qRoot
		uf.size[qRoot] += uf.size[pRoot]
	} else {
		uf.parent[qRoot] = pRoot
		uf.size[pRoot] += uf.size[qRoot]
	}

	uf.setCount--
}

func (uf *MyUnionFind) find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func TestMakeConnected(t *testing.T) {
	fmt.Println(makeConnected(12, [][]int{{1, 5}, {1, 7}, {1, 2}, {1, 4}, {3, 7}, {4, 7}, {3, 5}, {0, 6}, {0, 1}, {0, 4}, {2, 6}, {0, 3}, {0, 2}}))
}
