package didi_2023_03_12

import (
	"fmt"
	"math"
)

/*

一家新公司请你帮他设计Logo。该Logo比较特别，是一个尺寸为 2nX2n 的正方形图案，一些位置染上了黑色，其他位置则是白色。

Logo组成如下图：

其组成为：每次将该正方形的两条对称轴画出来，分割成 4 个尺寸相同的 2n-1 X 2n-1 的正方形。如果这是第奇数次分割，就把左上和右下的正方形全部涂黑，剩下两个正方形继续分割；
如果这是第偶数次分割，就把左下和右上的正方形全部涂黑，剩下两个正方形继续分割；
直到不能分割为止。
现在，输入 n，你需要输出这个Logo的图案。由于输出可能会很多，只需要输出特定的 q 行的内容即可。

第一行两个正整数 n,q；
接下来一行输入q个互不相同的正整数x1,x2,...,xq，表示需要依次输出该第x1行、第x2行,...，第xq行的内容。

对于所有的数据：
1≤n≤10，1≤q≤min(300,2n),1≤xi≤2n

似乎需要对二维矩阵进行递归

*/

func main() {
	var n, q int
	fmt.Scanln(&n, &q)
	lineNumber := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&lineNumber[i])
	}
	line := int(math.Pow(2, float64(n)))
	matrix := make([][]int, 0)
	for i := 0; i < line; i++ {
		matrix = append(matrix, make([]int, line))
	}
	// 初始都是0 表示白色  模拟涂黑色的过程
	// 用上下左右来固定边界

	var dfs func(arr [][]int, up int, down int, left int, right int, depth int)
	dfs = func(arr [][]int, up int, down int, left int, right int, depth int) {
		if depth > n {
			return
		}
		// depth 为奇数 处理 捺  depth为偶数 处理 撇

		rowMid := (up + down) / 2
		colMid := (left + right) / 2

		if depth%2 == 1 {

			for i := up; i <= rowMid; i++ {
				for j := left; j <= colMid; j++ {
					matrix[i][j] = 1
				}
			}

			for i := rowMid + 1; i <= down; i++ {
				for j := colMid + 1; j <= right; j++ {
					matrix[i][j] = 1
				}
			}

			dfs(matrix, up, rowMid, colMid+1, right, depth+1)
			dfs(matrix, rowMid+1, down, left, colMid, depth+1)

		} else {
			for i := up; i <= rowMid; i++ {
				for j := colMid + 1; j <= right; j++ {
					matrix[i][j] = 1
				}
			}
			for i := rowMid + 1; i <= down; i++ {
				for j := left; j <= colMid; j++ {
					matrix[i][j] = 1
				}
			}

			dfs(matrix, up, rowMid, left, colMid, depth+1)
			dfs(matrix, rowMid+1, down, colMid+1, right, depth+1)

		}

	}

	dfs(matrix, 0, len(matrix)-1, 0, len(matrix)-1, 1)
	for i := 0; i < len(lineNumber); i++ {
		rowNumber := lineNumber[i] - 1
		row := matrix[rowNumber]
		for j := 0; j < len(row); j++ {
			if row[j] == 0 {
				fmt.Print("W")
			} else {
				fmt.Print("B")
			}
		}
		fmt.Println()
	}
}
