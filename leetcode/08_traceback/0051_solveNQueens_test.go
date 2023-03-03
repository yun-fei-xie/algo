package _8_traceback

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/n-queens/description/

n层递归
*/

func solveNQueens(n int) [][]string {

	res := make([][]string, 0)
	position := make([]int, 0) // 存放路径

	row := make([]bool, n)
	col := make([]bool, n)
	dail1 := make([]bool, 2*n-1) // pie 2*n-1 而不是2*(n-1)
	dail2 := make([]bool, 2*n-1) // nai

	var putQueue func(index int) // index 表示行号
	putQueue = func(index int) {
		if index == n {
			res = append(res, generate(position, n))
			return
		}

		for i := 0; i < n; i++ { //i表示列号
			if !row[i] && !col[i] && !dail1[index+i] && !dail2[index-i+n-1] {
				row[i] = true
				col[i] = true
				dail1[index+i] = true
				dail2[index-i+n-1] = true

				position = append(position, i)
				putQueue(index + 1)
				position = position[:len(position)-1]

				row[i] = false
				col[i] = false
				dail1[index+i] = false
				dail2[index-i+n-1] = false
			}
		}
	}

	putQueue(0)

	return res
}
func generate(position []int, n int) []string {
	res := make([]string, 0)
	for i := 0; i < len(position); i++ {
		sb := strings.Builder{}
		for j := 0; j < n; j++ {
			if j != position[i] {
				sb.WriteString(".")
			} else {
				sb.WriteString("Q")
			}
		}
		res = append(res, sb.String())
	}
	return res
}

func TestNQueens(t *testing.T) {

	n := 4
	res := solveNQueens(n)
	for i := 0; i < len(res); i++ {

		fmt.Println(res[i])
	}

}
