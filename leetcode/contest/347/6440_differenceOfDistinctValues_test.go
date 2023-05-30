package _347

import (
	"fmt"
	"testing"
)

func differenceOfDistinctValues(grid [][]int) [][]int {
	row, col := len(grid), len(grid[0])
	answer := make([][]int, row)
	for i := 0; i < row; i++ {
		answer[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			answer[i][j] = Helper(grid, i, j, row, col)
		}
	}
	return answer
}

func Helper(grid [][]int, i, j, row, col int) int {

	// 用2个map分别记录左边有多少数字 右边
	leftMap := make(map[int]struct{})
	for leftI, leftJ := i, j; ; {
		if leftI-1 >= 0 && leftJ-1 >= 0 {
			if _, found := leftMap[grid[leftI-1][leftJ-1]]; !found {
				leftMap[grid[leftI-1][leftJ-1]] = struct{}{}
			}
			leftI--
			leftJ--
		} else {
			break
		}
	}

	rightMap := make(map[int]struct{})
	for rightI, rightJ := i, j; ; {
		if rightI+1 < row && rightJ+1 < col {
			if _, found := rightMap[grid[rightI+1][rightJ+1]]; !found {
				rightMap[grid[rightI+1][rightJ+1]] = struct{}{}
			}
			rightI++
			rightJ++
		} else {
			break
		}
	}
	leftCount := len(leftMap)
	rightCount := len(rightMap)
	if leftCount > rightCount {
		return leftCount - rightCount
	}
	return rightCount - leftCount
}

func TestDifferenceOfDistinctValues(t *testing.T) {
	fmt.Println(differenceOfDistinctValues([][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}))
}
