package easy

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/counting-bits/description/
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

计算bit 1的位数，使用Brian Kernighan 算法
或者是直接使用内置函数


*/

func countBits(n int) []int {
	res := make([]int, 0)
	for i := 0; i <= n; i++ {
		res = append(res, counts1(i))
		//	bits.OnesCount(uint(i))
	}
	return res
}

func counts1(x int) int {
	ans := 0
	for x > 0 {
		x = x & (x - 1)
		ans++
	}
	return ans
}

func TestCountBits(t *testing.T) {
	res := countBits(5)
	fmt.Println(res)
}
