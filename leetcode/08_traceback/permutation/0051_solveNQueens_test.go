package permutation

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
	var ans [][]string
	var position []int
	row := make([]int, n)
	col := make([]int, n)
	pie := make([]int, 2*n-1)
	na := make([]int, 2*n-1)

	var putQueue func(depth int)
	putQueue = func(depth int) {
		if depth >= n {
			solve := generate(position, n)
			ans = append(ans, solve)
			return
		}

		for i := 0; i < n; i++ { // 尝试在depth这一行上放置一枚皇后，坐标为（depth , i）
			if row[depth] == 0 && col[i] == 0 && pie[depth+i] == 0 && na[depth-i+n-1] == 0 { //可以放
				row[depth] = 1
				col[i] = 1
				pie[depth+i] = 1
				na[depth-i+n-1] = 1

				position = append(position, i)
				putQueue(depth + 1)
				position = position[:len(position)-1]

				row[depth] = 0
				col[i] = 0
				pie[depth+i] = 0
				na[depth-i+n-1] = 0
			}
		}
	}
	putQueue(0)
	return ans

}
func generate(position []int, n int) []string {
	var res []string
	for i, l := 0, len(position); i < l; i++ {
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
