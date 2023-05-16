package _0_dp

import (
	"fmt"
	"testing"
)

/*
70. 爬楼梯
https://leetcode.cn/problems/climbing-stairs/description/

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

方法：本质是一个求解斐波那契数列的过程。
递归: 树结构 开始有n级，每次1级或者2级，F(n) =  F(n-1) + F(n-2)

	递归终止的时候，n==1 只有一种爬法 return 1
				n==2 有2种爬法 return 2
*/
func climbStairs1(n int) int {

	var dfs func(stairs int) int
	dfs = func(stairs int) int {
		if stairs == 1 || stairs == 2 {
			return stairs
		} else {
			return dfs(stairs-1) + dfs(stairs-2)
		}
	}
	return dfs(n)
}

/*
记忆化搜索
*/
func climbStairs2(n int) int {
	mem := make([]int, n)

	var dfs func(stairs int) int
	dfs = func(stairs int) int {
		if stairs == 1 || stairs == 2 {
			return stairs
		}
		if mem[stairs-1] == 0 { //记忆化搜索
			mem[stairs-1] = dfs(stairs - 1)
		}
		if mem[stairs-2] == 0 { // 记忆化搜索
			mem[stairs-2] = dfs(stairs - 2)
		}

		return mem[stairs-1] + mem[stairs-2]
	}
	return dfs(n)
}

/*
递推方式|动态规划
*/
func climbStairs3(n int) int {
	if n <= 2 {
		return n
	}
	mem := make([]int, n+1)
	mem[1] = 1
	mem[2] = 2

	for i := 3; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}

func TestClimbStairs(t *testing.T) {

	res1 := climbStairs1(3)
	fmt.Println(res1)

	res2 := climbStairs2(3)
	fmt.Println(res2)

	res3 := climbStairs3(3)
	fmt.Println(res3)
}
