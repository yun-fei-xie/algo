package subset

import (
	"fmt"
	"math/bits"
	"testing"
)

/*
2397. 被列覆盖的最多行数
https://leetcode.cn/problems/maximum-rows-covered-by-columns/

给你一个下标从 0 开始的 m x n 二进制矩阵 mat 和一个整数 cols ，表示你需要选出的列数。
如果一行中，所有的 1 都被你选中的列所覆盖，那么我们称这一行 被覆盖 了。
请你返回在选择 cols 列的情况下，被覆盖 的行数 最大 为多少。

方法：枚举
如何快速判断，枚举之后的结果，能够覆盖多少行的一。
可以用一个数组记录每一行1的个数。


方法2：二进制枚举
对于二进制数组，可以使用二进制枚举的方式高效解题。


*/

func maximumRows(matrix [][]int, numSelect int) int {
	var row = len(matrix)
	var col = len(matrix[0])
	// 将二进制数组转化为数字 使用移位操作
	var mask = make([]int, row)
	for i := 0; i < row; i++ {
		sum := 0
		for j := 0; j < col; j++ {
			sum += matrix[i][j] << (col - 1 - j)
		}
		mask[i] = sum
	}

	var ans int
	//{0,1,0...} -> [0...2^n-1] 所有的可能性
	for i := 0; i < 2<<col; i++ {
		var cnt = 0
		if bits.OnesCount(uint(i)) == numSelect {
			for j := 0; j < row; j++ {
				if i&mask[j] == mask[j] {
					cnt++
				}
			}
			if cnt > ans {
				ans = cnt
			}

		}
	}

	return ans
}

func TestMaximumRows(t *testing.T) {
	//fmt.Println(bits.OnesCount(7)) // 7->111
	//fmt.Println(maximumRows([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}, 2))
	fmt.Println(maximumRows([][]int{{0, 0, 1, 0, 1}, {0, 1, 0, 0, 1}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 1, 0, 0}}, 1))
	//fmt.Println(2 << 2)
}
