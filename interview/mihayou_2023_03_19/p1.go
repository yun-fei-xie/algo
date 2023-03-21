package main

import (
	"fmt"
)

/*
BG被认为是一种颜色时有多少连通分量 -> 3
BG被认为是二种颜色时有多少连通分量 -> 6

输入
2 6
RRGGBB
RGBGRR

输出
3

如何表示这个图
*/
func main1() {
	var row, col int
	fmt.Scan(&row, &col)
	matrix := make([][]uint8, 0)
	for i := 0; i < row; i++ {
		var str string
		fmt.Scanln(&str)
		color := make([]uint8, len(str))
		for j := 0; j < len(str); j++ {
			color[j] = str[j]
		}
		matrix = append(matrix, color)
	}

	cnt1 := calc(&matrix, row, col, true)
	cnt2 := calc(&matrix, row, col, false)

	fmt.Println(cnt1 - cnt2)

}

func calc(graph *[][]uint8, row int, col int, flag bool) int {

	visit := make([][]bool, 0)
	for i := 0; i < row; i++ {
		visit = append(visit, make([]bool, col))
	}

	cnt := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			if visit[i][j] == false {
				cnt++
				if flag {
					dfs(graph, &visit, i, j)
				} else {
					dfs2(graph, &visit, i, j)
				}
			}
		}
	}
	return cnt
}

func dfs2(graph *[][]uint8, visit *[][]bool, i int, j int) {
	// 首先访问当前点
	(*visit)[i][j] = true

	// 左
	if j-1 >= 0 && (*visit)[i][j-1] == false {
		if (*graph)[i][j-1] == (*graph)[i][j] || ((*graph)[i][j] == 'B' && (*graph)[i][j-1] == 'G') || ((*graph)[i][j] == 'G' && (*graph)[i][j-1] == 'B') {
			dfs2(graph, visit, i, j-1)
		}
	}
	if j+1 < len((*graph)[0]) && (*visit)[i][j+1] == false {
		if (*graph)[i][j+1] == (*graph)[i][j] || ((*graph)[i][j] == 'B' && (*graph)[i][j+1] == 'G') || ((*graph)[i][j] == 'G' && (*graph)[i][j+1] == 'B') {
			dfs2(graph, visit, i, j+1)
		}
	}

	if i-1 >= 0 && (*visit)[i-1][j] == false {
		if (*graph)[i][j] == (*graph)[i-1][j] || ((*graph)[i][j] == 'B' && (*graph)[i-1][j] == 'G') || ((*graph)[i][j] == 'G' && (*graph)[i-1][j] == 'B') {
			dfs2(graph, visit, i-1, j)
		}
	}
	if i+1 < len(*graph) && (*visit)[i+1][j] == false {
		if (*graph)[i][j] == (*graph)[i+1][j] || ((*graph)[i][j] == 'B' && (*graph)[i+1][j] == 'G') || ((*graph)[i][j] == 'G' && (*graph)[i+1][j] == 'B') {
			dfs2(graph, visit, i+1, j)
		}
	}

}

func dfs(graph *[][]uint8, visit *[][]bool, i int, j int) {
	// 首先访问当前点
	(*visit)[i][j] = true

	// 左
	if j-1 >= 0 && (*graph)[i][j-1] == (*graph)[i][j] && (*visit)[i][j-1] == false {
		dfs(graph, visit, i, j-1)
	}

	// 右
	if j+1 < len((*graph)[0]) && (*graph)[i][j] == (*graph)[i][j+1] && (*visit)[i][j+1] == false {
		dfs(graph, visit, i, j+1)
	}

	// 上

	if i-1 >= 0 && (*graph)[i-1][j] == (*graph)[i][j] && (*visit)[i-1][j] == false {
		dfs(graph, visit, i-1, j)
	}

	// 下
	if i+1 < len(*graph) && (*graph)[i][j] == (*graph)[i+1][j] && (*visit)[i+1][j] == false {
		dfs(graph, visit, i+1, j)
	}
}
