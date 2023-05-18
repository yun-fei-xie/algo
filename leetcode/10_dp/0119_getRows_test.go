package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/pascals-triangle-ii/description/
119. 杨辉三角 II
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。
1
1 1
1 2 1
1 3 3 1
....
倒着计算，一个数组也可以搞定
*/
func getRow(rowIndex int) []int {

	arr1 := make([]int, rowIndex+1)
	arr1[0] = 1
	if rowIndex == 0 {
		return arr1
	}

	arr2 := make([]int, rowIndex+1)
	for i := 1; i <= rowIndex; i++ {
		// 每次迭代更新i个数字
		for j := 0; j <= i; j++ {
			var prev int
			if j != 0 {
				prev = arr1[j-1]
			}
			arr2[j] = arr1[j] + prev
		}
		// 每次都让arr1成为答案
		arr1, arr2 = arr2, arr1
	}
	return arr1
}

func TestGetRow(t *testing.T) {
	fmt.Println(getRow(3))
}
