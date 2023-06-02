package unionFind_test

import (
	"fmt"
	"testing"
)

/*
959. 由斜杠划分区域
https://leetcode.cn/problems/regions-cut-by-slashes/

在由 1 x 1 方格组成的 n x n 网格 grid 中，每个 1 x 1 方块由 '/'、'\' 或空格构成。这些字符会将方块划分为一些共边的区域。
给定网格 grid 表示为一个字符串数组，返回 区域的数量 。
请注意，反斜杠字符是转义的，因此 '\' 用 '\\' 表示。

方法：并查集
 1. \\的长度其实是1而不是2
 2. 为了应对\和/对一个单元格进行划分，这个题需要将单元格拆分成粒度更加小的格子。也就是拆成4个小的三角形。
 3. 并查集的编号，将二维的网格展开成一维，粒度到三角形。如果某个格子的坐标是(i,j)，那么这个格子前面有i行格子(因为下标从0开始)
    在这一行，它的前面还有j个格子。所以这个格子前面有i*n*4+j*4个三角形。对于每个格子，按照顺序对格子中的三角形进行编号即可。
 4. 最后的答案就是整个并查集的连通分量的个数。
*/
func regionsBySlashes(grid []string) int {
	n := len(grid)
	uf := newRegionUf(n * n * 4)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ch := grid[i][j]
			// 求出前面的序号 i*n*4+j*4
			// 如果当前坐标为（i,j）那么当前格子前面就有i行+这一行的j个格子（因为下标是从0开始的）
			pre := i*n*4 + j*4

			if ch == '/' {

				uf.unionElements(pre+0, pre+1)
				uf.unionElements(pre+2, pre+3)

			} else if ch == '\\' {
				uf.unionElements(pre+0, pre+3)
				uf.unionElements(pre+1, pre+2)

			} else {
				uf.unionElements(pre+0, pre+1)
				uf.unionElements(pre+1, pre+2)
				uf.unionElements(pre+2, pre+3)
			}
			// 单元格之间
			if j+1 < n {
				uf.unionElements(pre+3, pre+4+1)
			}
			if i+1 < n {
				uf.unionElements(pre+2, (i+1)*n*4+j*4)
			}
		}
	}
	return uf.sizeCount
}

type regionUf struct {
	parent    []int
	sizeCount int
}

func newRegionUf(n int) *regionUf {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &regionUf{
		parent:    p,
		sizeCount: n,
	}
}
func (uf *regionUf) find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}
func (uf *regionUf) unionElements(p int, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	uf.sizeCount--
	uf.parent[pRoot] = qRoot
}

func TestRegionsBySlashes(t *testing.T) {
	//fmt.Println(len("\\")) // len("\\") -> 1
	fmt.Println(regionsBySlashes([]string{" /", "/ "}))
}
