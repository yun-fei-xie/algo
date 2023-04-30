package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/triangle/
*/

/*
递归解法（测试超时)
难点在于在这样一个二维数组中进行移动。
仔细观察后发现，当前行（x , y）向下一步可以到（x+1，y）和(x+1,y+1)
*/
func minimumTotal(triangle [][]int) int {
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	var dfs func(row int, col int, depth int) (m int)
	dfs = func(row int, col int, depth int) (m int) {
		if depth == len(triangle)-1 {
			return triangle[row][col]
		}

		return min(dfs(row+1, col, depth+1), dfs(row+1, col+1, depth+1)) + triangle[row][col]
	}

	return dfs(0, 0, 0)
}

/*
记忆化搜索(哪里会重复)

		  2
		 3 4
		6 5 7
	   4 1 8 3

比如5这个位置的最小值，会同时被3和4用到，但是纯递归会重复计算
仔细观察递归的过程中那个min函数
*/
func minimumTotal2(triangle [][]int) int {
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	//记忆最小值
	mem := make([][]int, 0)
	for i := 0; i < len(triangle)-1; i++ {
		mem = append(mem, make([]int, len(triangle[i])))
	}

	var dfs func(row int, col int, depth int) (m int)
	dfs = func(row int, col int, depth int) (m int) {
		if depth == len(triangle)-1 {
			return triangle[row][col]
		}

		if mem[row][col] == 0 {
			mem[row][col] = min(dfs(row+1, col, depth+1), dfs(row+1, col+1, depth+1)) + triangle[row][col]
		}
		return mem[row][col]

	}

	return dfs(0, 0, 0)
}

/*
动态规划
从最低层开始思考： opt[row][col] = Min( opt[row][col] , opt[row+1][col])+ triangle[row][col]
2
3 4
6 5 7
4 1 8 3

数组的宽度最宽为 l = len(triangle[len(triangle)-1]) 我用这样一个一维数组dp[l]保存最小值状态，
层层向上，通过覆盖的方式最后返回dp[0] 则得到最小值
把这个dp数组开成二维空间其实比价容易理解，这里写成一维空间为了压缩空间

*/

func minimumTotal3(triangle [][]int) int {
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	dp := make([]int, len(triangle[len(triangle)-1])+1) //在末尾多放一个位置，这样计算三角形右下角元素的值的时候，不会产生数组越界。

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

/*
dp开成二维空间 和 triangle同形状，然后倒着计算,这个和递归完全反过来
*/
func minimumTotal4(triangle [][]int) int {
	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	dp := make([][]int, 0)
	for i := 0; i < len(triangle); i++ {
		dp = append(dp, make([]int, len(triangle[i])))
	}

	for j := 0; j < len(triangle[len(triangle)-1]); j++ {
		dp[len(triangle)-1][j] = triangle[len(triangle)-1][j] // 最后一行的最小值就是它自己，直接赋值过来，这是初始条件
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res := minimumTotal(triangle)
	fmt.Println(res)

	res2 := minimumTotal2(triangle)
	fmt.Println(res2)

	res3 := minimumTotal3(triangle)
	fmt.Println(res3)

	res4 := minimumTotal4(triangle)
	fmt.Println(res4)
}
