package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/unique-paths/description/

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？

动态规划 递推公式比较简单，mem[i][j] = mem[i+1][j] + mem[i][j+1]

如果递归的话怎么做?
递归的终止条件是什么？ 到达终点，或者越界（这个时候也要回退）
因此最后一行总是可以向右走，最右一列总是可以向下走

考虑到这些校验条件，我使用dp一开始就避开了这些检查
*/

/*
使用递归搜索
起点（0，0） 终点（m-1，n-1）
*/
func uniquePathsRec(m int, n int) int {
	var ans int
	var traceback func(x, y int)
	traceback = func(x, y int) {
		if x == m-1 && y == n-1 {
			ans++
			return
		}
		// 越界
		if x > m-1 || y > n-1 {
			return
		}
		// traceback(x+1,y） 和 traceback(x, y+1)会产生重叠子问题
		traceback(x+1, y) //向下
		traceback(x, y+1) //向右
	}
	traceback(0, 0)
	return ans
}

/*
改造上面的深度优先搜索，
*/

func uniquePathsRec2(m int, n int) int {
	// 二维记忆数组
	//mem := make([][]int, 0)
	//for i := 0; i <= m; i++ {
	//	mem = append(mem, make([]int, n))
	//}
	//mem[m-1][n-1] = 1

	mem := make([][]int, 0)
	for i := 0; i < m; i++ {
		mem = append(mem, make([]int, n))
	}

	var traceback func(x, y int) int
	traceback = func(x, y int) int {
		if x == m-1 && y == n-1 {
			return 1
		}
		// 越界
		if x > m-1 || y > n-1 {
			return 0
		}
		if mem[x+1][y] == 0 {
			mem[x+1][y] = traceback(x+1, y)
		}
		if mem[x][y+1] == 0 {
			mem[x][y+1] = traceback(x, y+1)
		}

		return mem[x+1][y] + mem[x][y+1]
	}
	return traceback(0, 0)

}

func uniquePaths(m int, n int) int {

	mem := make([][]int, 0)
	for i := 0; i < m; i++ {
		mem = append(mem, make([]int, n))
	}

	// 初始条件 最下面一行，只有一种走法
	for i := 0; i < n; i++ {
		mem[m-1][i] = 1
	}
	// 最右边一列，只有一种走法
	for j := 0; j < m; j++ {
		mem[j][n-1] = 1
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			mem[i][j] = mem[i+1][j] + mem[i][j+1]
		}
	}

	return mem[0][0]

}

func TestUniquePath(t *testing.T) {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePathsRec(3, 7))
	fmt.Println(uniquePathsRec2(3, 7))
}
