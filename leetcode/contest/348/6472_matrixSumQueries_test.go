package _348__test

import (
	"fmt"
	"testing"
)

/*
2718. 查询后矩阵的和
https://leetcode.cn/problems/sum-of-matrix-after-queries/

给你一个整数 n 和一个下标从 0 开始的 二维数组 queries ，其中 queries[i] = [typei, indexi, vali] 。
一开始，给你一个下标从 0 开始的 n x n 矩阵，所有元素均为 0 。每一个查询，你需要执行以下操作之一：

如果 typei == 0 ，将第 indexi 行的元素全部修改为 vali ，覆盖任何之前的值。
如果 typei == 1 ，将第 indexi 列的元素全部修改为 vali ，覆盖任何之前的值。
请你执行完所有查询以后，返回矩阵中所有整数的和。

暴力方法：空间复杂度爆炸

有没有办法知道，每次query操作会为矩阵增加多少值
比赛的时候，思路是对的，但是没有去逆序思考

方法：倒序遍历queries
倒着遍历 hash表记录每一行 每一列之前是否已经被覆盖
如果是行，那么如果这一行之前没有被覆盖过（倒序）则统计这一行中有多少列之前被操作过。
如果是列，那么如果这一列之前没有被覆盖过（倒序）则统计这一列中有多少行之前被操作过。
相应地可以计算出每次修改可以拿到的贡献值。
*/
func matrixSumQueries(n int, queries [][]int) int64 {

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for _, query := range queries {
		typ := query[0]
		index := query[1]
		val := query[2]
		if typ == 0 {
			for i := 0; i < n; i++ {
				matrix[index][i] = val
			}
		}
		if typ == 1 {
			for i := 0; i < n; i++ {
				matrix[i][index] = val
			}
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += matrix[i][j]
		}
	}
	return int64(ans)
}

/*
倒着遍历 hash表记录每一行 每一列之前是否已经被覆盖
如果是行，那么如果这一行之前没有被覆盖过（倒序）则统计这一行中有多少列之前被操作过。
如果是列，那么如果这一列之前没有被覆盖过（倒序）则统计这一列中有多少行之前被操作过。
相应地可以计算出每次修改可以拿到的贡献值。
*/

func matrixSumQueries2(n int, queries [][]int) int64 {

	row := make(map[int]struct{})
	col := make(map[int]struct{})

	var ans int
	for i := len(queries) - 1; i >= 0; i-- {
		tp := queries[i][0]
		index := queries[i][1]
		val := queries[i][2]
		// 如果是行 看看之前已经有多少列被覆盖(并且这一行之前没有被操作过)
		if tp == 0 {
			if _, found := row[index]; !found {

				coverCol := len(col)
				ans += (n - coverCol) * val
				row[index] = struct{}{}
			}
		} else {
			if _, found := col[index]; !found {

				coverRow := len(row)
				ans += (n - coverRow) * val
				col[index] = struct{}{}
			}
		}
	}
	return int64(ans)
}

func matrixSumQueries3(n int, queries [][]int) int64 {

	row := make(map[int]bool)
	col := make(map[int]bool)
	visited := []map[int]bool{row, col}

	var ans int64
	for i := len(queries) - 1; i >= 0; i-- {
		tp, index, val := queries[i][0], queries[i][1], queries[i][2]
		if !visited[tp][index] {
			ans += int64((n - len(visited[1-tp])) * val)
			// 1-tp -> tp^1 实现0与1的互换
			visited[tp][index] = true
		}
	}
	return ans
}

func TestMatrixSumQuery(t *testing.T) {
	fmt.Println(matrixSumQueries3(3, [][]int{{0, 0, 1}, {1, 2, 2}, {0, 2, 3}, {1, 0, 4}}))
}
