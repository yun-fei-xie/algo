package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/integer-break/description/

当没有头绪的时候要想一想，如果使用暴力解法 枚举 应该怎么做？
在这个题目里面，因为不知道n有多大，因此无法知道需要使用几重循环,那么可以考虑使用递归。

怎么分呢？比如n=4 (3+? 2+? 1+? ) ？号处再次进行递归
*/
func integerBreak(n int) int {
	mem := make([]int, n+1) // 记忆化搜索
	// 递归终止条件 n==1
	var dfs func(num int) int

	dfs = func(num int) int {
		if num == 1 {
			return 1
		}
		if mem[num] != 0 {
			return mem[num]
		}
		for i := 1; i < num; i++ {
			mem[num] = max3(mem[num], i*(num-i), i*dfs(num-i)) // 这个地方有一个陷阱，比大小的时候加上i*(num-i) 向下递归的时候其实就相当于在mem数组中寻找值
		}
		return mem[num]
	}
	return dfs(n)
}

func max3(i, j, k int) int {
	var temp int
	if i > j {
		temp = i
	} else {
		temp = j
	}
	if k > temp {
		return k
	}
	return temp
}

/*
动态规划:只要能递归+记忆化 就可以反着推倒
*/

func integerBreak2(n int) int {
	mem := make([]int, n+1) // mem[i] -> i经过分割后的最大值
	mem[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			// 两个部分j*(i-j)  或者继续分割（i-j）
			mem[i] = max3(mem[i], j*(i-j), mem[i-j])
		}
	}
	return mem[n]
}

func TestIntegerBreak(t *testing.T) {
	fmt.Println(integerBreak(10))
	fmt.Println(integerBreak2(10))

}
