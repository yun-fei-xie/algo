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

func uniquePaths(m int, n int) int {

	mem := make([][]int, 0)
	for i := 0; i < m; i++ {
		mem = append(mem, make([]int, n))
	}

	// 初始条件
	for i := 0; i < n; i++ {
		mem[m-1][i] = 1
	}
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
}
