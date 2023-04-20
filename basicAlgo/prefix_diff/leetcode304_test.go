package prefix_diff

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/range-sum-query-2d-immutable/description/
给定一个二维矩阵 matrix，以下类型的多个请求：
计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：
NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。

二维矩阵的前缀和与差分

1 2 4 3
5 1 2 4 -> 3 * 4 的矩阵
6 3 5 9

前缀和：
1  3  7  10
6  9  15 22
12 18 29 45

在原始矩阵上多一行和多一列
0 0 0 0 0
0 1 2 4 3
0 5 1 2 4 -> 4 * 5 的矩阵 原来的坐标整体平移一位 （x ,y ） -> (x+1 , y+1)
0 6 3 5 9

[0 0 0 0 0]
[0 1 3 7 10]
[0 6 9 15 22]
[0 12 18 29 45]


*/

type NumMatrix struct {
	prefixMatrix [][]int
}

func ConstructorNumMatrix(matrix [][]int) NumMatrix {
	row := len(matrix)
	col := len(matrix[0])
	prefix := make([][]int, 0)
	for i := 0; i <= row; i++ {
		prefix = append(prefix, make([]int, col+1))
	}
	// 最上方多放一排，最左边多放一列。后面的处理，所有的坐标整体都需要+1，进行平移。
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{prefixMatrix: prefix}
}

/*
将数组扩充之后，不需要再对边界进行讨论
*/
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	return this.prefixMatrix[row2+1][col2+1] - this.prefixMatrix[row1][col2+1] - this.prefixMatrix[row2+1][col1] + this.prefixMatrix[row1][col1]

}

func TestNumMatrix(t *testing.T) {
	var matrix = [][]int{{1, 2, 4, 3}, {5, 1, 2, 4}, {6, 3, 5, 9}}
	numMatrix := ConstructorNumMatrix(matrix)
	for i := 0; i < len(numMatrix.prefixMatrix); i++ {
		fmt.Println(numMatrix.prefixMatrix[i])
	}
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 *
 */
