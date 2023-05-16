package kanpsack

import (
	"fmt"
	"math"
	"testing"
)

/*
279. 完全平方数
https://leetcode.cn/problems/perfect-squares/description/

给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

方法：完全背包问题
每一个完全平方数都是一类物品。例如，1，4，9,16...
n就是背包的初始容量。 于是，题目就转化为，用最少数量的物品，恰好把背包装满。

*/

func numSquares(n int) int {
	// 这里人为构造了一个数组 可不可以去掉？
	squareNumbers := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		squareNumbers = append(squareNumbers, i*i)
	}

	// 这里的i表示物品的重量
	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 0
			}
			// 无法完成任务
			return math.MaxInt32
		}

		m := math.MaxInt32
		for j := 0; c-j*squareNumbers[i] >= 0; j++ {
			m = min(m, dfs(i-1, c-j*squareNumbers[i])+j)
		}
		return m
	}

	ans := dfs(len(squareNumbers)-1, n)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

/*
动态规划解法
*/
func numSquares2(n int) int {
	fn := make([]int, n+1)
	fn[0] = 0

	for i := 1; i <= n; i++ {
		min := math.MaxInt64
		for j := 1; j*j <= i; j++ { // 注意这里的遍历顺序，一定是从小到大（比如你在求fn(3) 你的fn(2)和fn(1)一定要在前面求过）
			if fn[i-j*j] < min {
				min = fn[i-j*j] // i-j*j <i 所以这个值一定在前面求解过
			}
		}
		fn[i] = min + 1
	}
	return fn[n]
}

func TestNumSquares(t *testing.T) {

	fmt.Println(numSquares(12))
	fmt.Println(numSquares2(12))

}
