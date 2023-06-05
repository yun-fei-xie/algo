package unionFind_test

import (
	"fmt"
	"testing"
)

/*
803. 打砖块
https://leetcode.cn/problems/bricks-falling-when-hit/description/

有一个 m x n 的二元网格 grid ，其中 1 表示砖块，0 表示空白。砖块 稳定（不会掉落）的前提是：
一块砖直接连接到网格的顶部，或者
至少有一块相邻（4 个方向之一）砖块 稳定 不会掉落时
给你一个数组 hits ，这是需要依次消除砖块的位置。每当消除 hits[i] = (rowi, coli) 位置上的砖块时，对应位置的砖块（若存在）会消失，然后其他的砖块可能因为这一消除操作而 掉落 。一旦砖块掉落，它会 立即 从网格 grid 中消失（即，它不会落在其他稳定的砖块上）。
返回一个数组 result ，其中 result[i] 表示第 i 次消除操作对应掉落的砖块数目。
注意，消除可能指向是没有砖块的空白位置，如果发生这种情况，则没有砖块掉落。

方法：
1. 需要注意的是，这个题中的稳定是一个**递归定义**。它表示的是砖块直接或者间接和第一行中的砖块相连。
2. 逆向思考，题目要求给出每次消除操作后，对应掉落的砖块的数目。可以考虑一开始把所有的需要敲掉的砖块在图中抹除。然后逐次添加，看看每次添加后，有多少个砖块会和天花板连上。
3. 第一行的砖块都是和天花板相连，因此可以考虑抽象出一个虚拟节点。这个节点代表天花板。（类似之前的图类型的题目中的超级节点）
*/

func hitBricks(grid [][]int, hits [][]int) []int {

	row, col := len(grid), len(grid[0])
	superNode := row * col
	uf := newBrickUF(row*col + 1)

	graphCopy := make([][]int, row)
	for i := 0; i < row; i++ {
		graphCopy[i] = make([]int, col)
		copy(graphCopy[i], grid[i])
	}
	//copy(graphCopy, grid) 直接对二维slice使用copy，只会copy一维slice的指针
	//所有，如果要使用deep copy的效果，
	// 把消除的点从图中先抹掉
	for _, hit := range hits {
		graphCopy[hit[0]][hit[1]] = 0
	}

	// 第一行的砖块和天花板进行连通
	for i := 0; i < col; i++ {
		if graphCopy[0][i] == 1 {
			uf.unionElements(i, superNode)
		}
	}
	// 处理剩下的行 按照每个顶点和上，左的砖块是否有联通进行判断
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if graphCopy[i][j] == 0 {
				continue
			}
			if graphCopy[i-1][j] == 1 {
				uf.unionElements((i-1)*col+j, i*col+j)
			}
			if j > 0 && graphCopy[i][j-1] == 1 {
				uf.unionElements(i*col+(j-1), i*col+j)
			}
		}
	}

	var ans = make([]int, len(hits))
	dirX := []int{-1, 1, 0, 0}
	dirY := []int{0, 0, -1, 1}
	// 倒着处理hits,看看以superNode为根节点的连通分量中的节点增加了多少
	for i := len(hits) - 1; i >= 0; i-- {
		x := hits[i][0]
		y := hits[i][1]
		if grid[x][y] == 0 {
			continue
		}
		// orig 应该放在这里
		orig := uf.getSize(superNode)

		if x == 0 {
			uf.unionElements(y, superNode)
		}

		//之前的size
		//	orig := uf.getSize(superNode)
		// 尝试把这个点和上下左右的点挂上关系
		for j := 0; j < 4; j++ {
			newX := x + dirX[j]
			newY := y + dirY[j]
			if inArea(newX, newY, row, col) && graphCopy[newX][newY] == 1 {
				uf.unionElements(x*col+y, newX*col+newY)
			}
		}
		curr := uf.getSize(superNode)

		if curr-orig > 1 {
			ans[i] = curr - orig - 1
		}
		graphCopy[x][y] = 1
	}
	return ans
}

func inArea(x, y, row, col int) bool {
	if x < 0 || x >= row || y < 0 || y >= col {
		return false
	}
	return true
}

type brickUF struct {
	parent []int
	size   []int
}

func newBrickUF(n int) *brickUF {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &brickUF{
		parent: p,
		size:   s,
	}
}

// getSize

func (uf *brickUF) getSize(p int) int {
	rootP := uf.find(p)
	return uf.size[rootP]
}

func (uf *brickUF) find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *brickUF) unionElements(p int, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
	uf.size[qRoot] += uf.size[pRoot]
}

func TestHitBricks(t *testing.T) {
	//fmt.Println(hitBricks([][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}}))
	fmt.Println(hitBricks([][]int{{1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0}, {0, 2}, {1, 1}}))
}
