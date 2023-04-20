package prefix_diff

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/maximal-square/
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

先想一下暴力解法怎么做

穷举所有可能的正方形。
遍历矩阵的所有的点，以该点为正方形的左上角，探索正方形（探索的过程中第一次遇到条件不成立就可以停下来，因为子矩形不成立，那么父正方形也不会成立）

*/

/*
暴力求解，效率低在valid 这里
*/
func maximalSquare(matrix [][]byte) int {

	var valid func(matrix [][]byte, x1, y1, x2, y2 int) bool
	valid = func(matrix [][]byte, x1, y1, x2, y2 int) bool {
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				if matrix[i][j] != '1' {
					return false
				}
			}
		}
		return true
	}

	var ans int

	var row = len(matrix)
	var col = len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 以 matrix[i][j]为顶点扩散
			var l int = 0
			for {
				if i+l < row && j+l < col {
					if valid(matrix, i, j, i+l, j+l) {

						if l+1 > ans {
							ans = l + 1
						}
						l++

					} else {
						break
					}
				} else {
					break
				}
			}

		}
	}

	return ans * ans
}

/*
前缀和解法：
相当于求解子矩阵和，等于边长的平方
*/
func maximalSquare2(matrix [][]byte) int {

	row := len(matrix)
	col := len(matrix[0])
	prefix := make([][]int, 0)
	for i := 0; i < row; i++ {
		prefix = append(prefix, make([]int, col))
	}

	// 前缀和
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			var prefix1, prefix2, prefix3 int
			if i-1 >= 0 {
				prefix1 = prefix[i-1][j]
			}

			if j-1 >= 0 {
				prefix2 = prefix[i][j-1]
			}

			if i-1 >= 0 && j-1 >= 0 {
				prefix3 = prefix[i-1][j-1]
			}
			var num int
			if matrix[i][j] == '1' {
				num = 1
			}
			prefix[i][j] = prefix1 + prefix2 - prefix3 + num
		}
	}

	var ans int

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			var l int

		loop:
			for {
				if i+l >= row || j+l >= col {
					break loop
				}

				if subMatrix(prefix, i, j, i+l, j+l) == (l+1)*(l+1) {
					if l+1 > ans {
						ans = l + 1
					}
					l++
				} else {
					break loop
				}
			}

		}
	}

	return ans * ans

}

func TestMaximalSquare(t *testing.T) {

	var matrix = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	res := maximalSquare2(matrix)
	fmt.Println(res)
}
