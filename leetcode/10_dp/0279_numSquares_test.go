package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/perfect-squares/description/

给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

*/

/*
 */

var mem []int

func numSquares(n int) int {
	mem = make([]int, n+1)
	for i := 0; i < len(mem); i++ {
		mem[i] = -1
	}
	return numSquaresDfs(n)
}

func numSquaresDfs(num int) int {
	if num == 0 {
		return 0
	}

	if mem[num] != -1 {
		return mem[num]
	}

	minHeight := math.MaxInt64
	for i := 1; i*i <= num; i++ {
		h := numSquaresDfs(num - i*i) // 重叠子问题
		if h < minHeight {
			minHeight = h
		}
	}

	mem[num] = minHeight + 1

	return mem[num]
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
