package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/unique-binary-search-trees/
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。


*/
/*
解法1：递归
*/

func numTrees1(n int) int {

	var traceback func(x, y int) int //[x...y] 节点x到节点y能够构成多少种搜索树
	traceback = func(x, y int) int {
		//只有一个节点
		if x >= y {
			return 1
		}

		// 寻找以k为根节点的搜索树的个数
		var sum int
		for k := x; k <= y; k++ {
			left := traceback(x, k-1)  //左子树的数量
			right := traceback(k+1, y) // 右子树的数量
			sum += left * right
		}
		return sum
	}
	return traceback(1, n)

}

/*
解法2：递归+记忆化搜索
*/
func numTrees2(n int) int {

	mem := make([]int, n+1) // 下标表示元素的个数[1...n]
	mem[1] = 1

	var search func(i, j int) int // search [i...j] 这个区间存在多少种BST
	search = func(i, j int) int {
		if i >= j { // 只有一种可能
			return 1
		}

		if mem[j-i+1] != 0 {
			return mem[j-i+1]
		}

		// 枚举以k（i<=k<=j）为根节点，以[i...j]为区间的BST的数量
		sum := 0
		for k := i; k <= j; k++ {

			left := search(i, k-1)
			right := search(k+1, j)
			sum += left * right
		}

		mem[j-i+1] = sum

		return mem[j-i+1]
	}
	return search(1, n)
}

/*
解法3：动态规划 能不能自底向上
*/

func numTrees3(n int) int {
	mem := make([]int, n+1) // 下标表示元素的个数[1...n]
	mem[1] = 1
	// todo 这个题不是很好想
	for i := 2; i <= n; i++ {

	}
	return -1
}

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees1(3))
	fmt.Println(numTrees2(3))
}
