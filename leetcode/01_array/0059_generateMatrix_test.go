package _1_array

import (
	"fmt"
	"testing"
)

/**
https://leetcode.cn/problems/spiral-matrix-ii/

第一反应：能不能模拟这个数字的填充轨迹或者计算出每个数字和下标之间的关系

1.模拟轨迹，需要遵循一定的原则 按照数字填充的顺序，分为4个步骤
使用两个指针定义每个步骤填充数字的左右边界



*/

func generateMatrix(n int) [][]int {
	// init
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	// 行的left、right 其实是列号
	// 列的left、right 其实是行号
	rowLeft := 0
	rowRight := n - 1
	colLeft := 0
	colRight := n - 1

	step := [4]int{1, 2, 3, 4} // 四步走战略
	stepIndex := 0

	for num := 1; num <= n*n; {
		if step[stepIndex%len(step)] == 1 {
			for i := colLeft; i <= colRight; i++ {
				res[rowLeft][i] = num // 行下标不变 移动列下标
				num++
			}
			rowLeft++
			stepIndex++
		} else if step[stepIndex%len(step)] == 2 {
			for i := rowLeft; i <= rowRight; i++ {
				res[i][colRight] = num // 列下标不变 移动行下标
				num++
			}
			colRight--
			stepIndex++
		} else if step[stepIndex%len(step)] == 3 {
			for i := colRight; i >= colLeft; i-- {
				res[rowRight][i] = num
				num++
			}
			rowRight--
			stepIndex++
		} else if step[stepIndex%len(step)] == 4 {
			for i := rowRight; i >= rowLeft; i-- {
				res[i][colLeft] = num
				num++
			}
			colLeft++
			stepIndex++
		}
	}
	return res
}

func TestGenerateMatrix(t *testing.T) {
	n := 3
	res := generateMatrix(n)

	for i := 0; i < n; i++ {
		fmt.Print("[ ")
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Printf("]\n")
	}
}
