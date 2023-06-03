package unionFind_test

import (
	"fmt"
	"testing"
)

/*
947. 移除最多的同行或同列石头
https://leetcode.cn/problems/most-stones-removed-with-same-row-or-column/description/

n 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头。
如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头。
给你一个长度为 n 的数组 stones ，其中 stones[i] = [xi, yi] 表示第 i 块石头的位置，返回 可以移除的石子 的最大数量。

题目容易理解错误，如果想要移除一个stone,那么这个stone的同行或者同列上必须要有其他stone。
一旦这个石头被移除，那么下一次搜索，这个石头就算没有了。

方法：并查集
1.建图涉及到给每一个点进行编码，不能简单横纵坐标相加，否则(1,0),(0,1)会相同。因为x的上界是10000，所以可以让横坐标乘以上界，再加上y。可以获得一个唯一的位置。
然后配合hash+自增id，给每个点进行编码。
2.假设一个连通分量有m个顶点，那么m-1个stone可以被移除。如果有k个顶点，r个连通分量，那么最后的结果就是k-r。
每一个连通分量都有一个点最后无法被移除。因为，前面的点被移除后，它找不到同行同列的stone。
*/
func removeStones(stones [][]int) int {
	n := len(stones)
	factor := 20000
	uf := newStoneUf(n)
	id := 0
	hashMap := make(map[int]int) // 根据坐标映射id
	for _, stone := range stones {
		pos := stone[0]*factor + stone[1]
		hashMap[pos] = id
		id++
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				idI := hashMap[stones[i][0]*factor+stones[i][1]]
				idJ := hashMap[stones[j][0]*factor+stones[j][1]]
				uf.unionElements(idI, idJ)
			}
		}
	}

	return n - uf.sizeCount

}

type stoneUf struct {
	parent    []int
	size      []int
	sizeCount int
}

func newStoneUf(n int) *stoneUf {
	p := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		size[i] = 1
	}
	return &stoneUf{
		parent:    p,
		size:      size,
		sizeCount: n,
	}
}

func (uf *stoneUf) find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *stoneUf) unionElements(p int, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
	uf.size[qRoot] += uf.size[pRoot]
	uf.sizeCount--
}

func TestRemoveStones(t *testing.T) {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
}
