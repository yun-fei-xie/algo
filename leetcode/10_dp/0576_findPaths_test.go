package _0_dp

import (
	"fmt"
	"testing"
)

/*
这个题和1575几乎一样

1. 越界（达到了目的地） 递归终止
2. 步数用完（递归终止）
3.步数没有用完 也没有越界 -> 尝试上下左右四个方向
*/
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var mod = 1000000007
	var ans int
	var dfs func(row int, col int, remainMove int)
	dfs = func(row int, col int, remainMove int) {
		// 找到一条路径
		if outOfBound(m, n, row, col) {
			ans++
			ans %= mod
			return
		}
		if remainMove == 0 {
			return
		}

		//上下左右
		dfs(row-1, col, remainMove-1)
		dfs(row+1, col, remainMove-1)
		dfs(row, col-1, remainMove-1)
		dfs(row, col+1, remainMove-1)
	}

	dfs(startRow, startColumn, maxMove)
	return ans
}

func outOfBound(m, n, x, y int) bool {
	if x < 0 || x >= m || y < 0 || y >= n {
		return true
	}
	return false
}

/*
修改成回溯
如果表达二维坐标、剩余最大步数 到达界外的数量，那么需要定义一个三维数组。
是否可以降维，将二维坐标降维到一维？（编码）
index = x * r + y;
(x, y) = (index / r, index % r);
r表示列号
3*5的矩阵，有15个元素。（0 ，1）-> 0 * 5 + 1 -> 1
(1,0) -> 1 * 5 + 1 = 6
*/
func findPaths2(m int, n int, maxMove int, startRow int, startColumn int) int {
	var mod = 1000000007
	var mem [][]int = make([][]int, m*n)
	for i := 0; i < m*n; i++ {
		k := make([]int, maxMove+1)
		for j := 0; j < maxMove+1; j++ {
			k[j] = -1
		}
		mem[i] = k
	}

	var dfs func(row, col, remainMove int) int // 返回从row , col 出发，最大步数remainMove
	dfs = func(row, col, remainMove int) int {
		// 当前正在界外
		if outOfBound(m, n, row, col) && remainMove >= 0 {
			//mem[row*r+col][remainMove] = 1
			return 1
		}

		if mem[row*n+col][remainMove] != -1 {
			return mem[row*n+col][remainMove]
		}

		if remainMove <= 0 {
			mem[row*n+col][remainMove] = 0
			return 0
		}

		var up = dfs(row-1, col, remainMove-1)
		var down = dfs(row+1, col, remainMove-1)
		var left = dfs(row, col-1, remainMove-1)
		var right = dfs(row, col+1, remainMove-1)

		var sum int
		sum += up + down + left + right
		sum %= mod
		mem[row*n+col][remainMove] = sum
		return sum
	}
	return dfs(startRow, startColumn, maxMove)
}

/*
将上面的递归+记忆化翻译成动态规划
*/

func TestFindPaths(t *testing.T) {
	//	fmt.Println(findPaths(2, 2, 2, 0, 0))
	//fmt.Println(findPaths2(2, 2, 2, 0, 0))
	fmt.Println(findPaths2(1, 3, 3, 0, 1))
}
