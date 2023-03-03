package _0_dp

import (
	"fmt"
	"testing"
)

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
网格中的障碍物和空位置分别用 1 和 0 来表示。

障碍物那一格的mem为0 表示此路不通

对于最下行与最右列 需要特殊处理（如果前方有障碍物，障碍物后面的部分肯定过不去）

*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	mem := make([][]int, 0)
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	for i := 0; i < len(obstacleGrid); i++ {
		mem = append(mem, make([]int, n))
	}

	hasObs := false
	for i := n - 1; i >= 0; i-- {
		if obstacleGrid[m-1][i] == 0 && hasObs == false { // 如果前方一直没有障碍物
			mem[m-1][i] = 1
			continue
		}
		//
		if obstacleGrid[m-1][i] == 1 {
			hasObs = true
		}
		mem[m-1][i] = 0
	}
	// 处理最右列
	hasObs = false

	for i := m - 1; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 0 && hasObs == false {
			mem[i][n-1] = 1
			continue
		}
		if obstacleGrid[i][n-1] == 1 {
			hasObs = true
		}
		mem[i][n-1] = 0
	}
	// 处理其他元素
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 { // 特殊处理
				mem[i][j] = 0
			} else {
				mem[i][j] = mem[i+1][j] + mem[i][j+1]
			}
		}
	}
	return mem[0][0]

}

func TestUniquePathWithOb(t *testing.T) {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
