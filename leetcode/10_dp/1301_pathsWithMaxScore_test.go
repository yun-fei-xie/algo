package _0_dp

import (
	"fmt"
	"testing"
)

/*
["E23",
 "2X2",
 "12S"]
*/
/*
如果可以用递归，那么问题就与规模无关。只是将问题的规模变小。
*/
func pathsWithMaxScore(board []string) []int {
	var mod int = 1e9 + 7
	m := len(board)
	n := len(board[0])

	var memPath = make([]int, m*n)
	var memScore = make([]int, m*n)
	for i := 0; i < m*n; i++ {
		memPath[i] = -1
	}
	//从row,col出发，达到左上角，返回路径数量和路径和
	var traceback func(board []string, row int, col int) (int, int)
	traceback = func(board []string, row int, col int) (int, int) {
		// 如果当前已经越界 越界永远放在前面进行检查
		if outOfBound(m, n, row, col) {
			return 0, 0
		}
		// 返回cache中的内容
		if memPath[row*n+col] != -1 {
			return memPath[row*n+col], memScore[row*n+col]
		}

		// 到达终点
		if row == 0 && col == 0 {
			return 1, 0
		}
		// 如果当前位置为障碍物
		if board[row][col] == 'X' {

			return 0, 0
		}
		// 上、右、左上
		p1, sum1 := traceback(board, row-1, col)
		p2, sum2 := traceback(board, row, col-1)
		p3, sum3 := traceback(board, row-1, col-1)

		max := max3(sum1, sum2, sum3)
		var path int
		if p1 != 0 && sum1 == max {
			path += p1
		}

		if p2 != 0 && sum2 == max {
			path += p2

		}
		if p3 != 0 && sum3 == max {
			path += p3
		}
		//如果当前是右下角
		if row == m-1 && col == n-1 {
			if path == 0 {
				return 0, 0
			}
			memPath[row*n+col] = path % mod
			memScore[row*n+col] = max % mod
			return memPath[row*n+col], memScore[row*n+col]
		} else { // 当前是一个合法区域
			if path == 0 {
				memPath[row*n+col] = 0
				memScore[row*n+col] = 0
				return 0, 0
			}
			memPath[row*n+col] = path % mod
			memScore[row*n+col] = (max + int(board[row][col]-'0')) % mod
			return memPath[row*n+col], memScore[row*n+col]
		}
	}

	path, maxScore := traceback(board, m-1, n-1)
	return []int{maxScore, path}
}

/*
如果时候记忆化？ 递归函数返回两个值，用两个一维数组。同时对二维坐标进行一维编码。
*/

func TestPathsWithMaxSocre(t *testing.T) {
	fmt.Println(pathsWithMaxScore([]string{"E23", "2X2", "12S"}))
	fmt.Println(pathsWithMaxScore([]string{"E12", "1X1", "21S"}))
	fmt.Println(pathsWithMaxScore([]string{"E11", "XXX", "11S"}))

}
