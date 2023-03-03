package _1_array

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/spiral-matrix/description/
*/
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)

	rowUp := 0
	rowDown := len(matrix) - 1
	colLeft := 0
	colRight := len(matrix[0]) - 1

	step := []int{1, 2, 3, 4}
	stepIndex := 0

	for i := 0; i < len(matrix)*len(matrix[0]); { // 数字的个数  注意 不是n*n 而是m*n

		if step[(stepIndex+len(step))%len(step)] == 1 {
			for j := colLeft; j <= colRight; j++ {
				res = append(res, matrix[rowUp][j])
				i++
			}
			rowUp++
			stepIndex++
		} else if step[(stepIndex+len(step))%len(step)] == 2 {
			for j := rowUp; j <= rowDown; j++ {
				res = append(res, matrix[j][colRight])
				i++
			}
			colRight--
			stepIndex++
		} else if step[(stepIndex+len(step))%len(step)] == 3 {
			for j := colRight; j >= colLeft; j-- {
				res = append(res, matrix[rowDown][j])
				i++
			}
			rowDown--
			stepIndex++
		} else if step[(stepIndex+len(step))%len(step)] == 4 {
			for j := rowDown; j >= rowUp; j-- {
				res = append(res, matrix[j][colLeft])
				i++
			}
			colLeft++
			stepIndex++
		}

	}

	return res

}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	res := spiralOrder(matrix)
	fmt.Println(res)

}
