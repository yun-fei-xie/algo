package mid

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/unique-binary-search-trees/description/?favorite=2cktkvj
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。



这个题需要一定的分析能力。

假设 n 个节点存在二叉排序树的个数是 G (n)，令 f(i) 为以 i 为根的二叉搜索树的个数，则
G(n)=f(1)+f(2)+f(3)+f(4)+...+f(n)

当 i 为根节点时，其左子树节点个数为 i-1 个，右子树节点为 n-i，则

f(i)=G(i−1)∗G(n−i)  这里是乘积的关系（左边m种情况，右边n种情况 那么结合起来就是m*n种情况）

两个公式结合，得到了递归关系。



递归分析：核心分析点在于对根节点进行枚举。然后将区间划为2个部分，左边部分能够形成的bst数量 乘以 右边部分能够形成的bst数量
左右区间可以继续递归，直到某个区间只有一个节点（只有一种可能->直接返回）

假设n=3 , 那么会有3种情况，以1为根节点，以2为根节点，以3为根节点。
当以1为根节点的话，由于是二叉搜索树，所以节点2、3都在根节点的右子树上，因此左子树有0个节点、右子树有2个节点。
再递归到右子树上，[2,3]两个节点，以2为右子树的根节点、以3为右子树的根节点。
在划分的过程中会出现区间重叠问题，因此可以使用记忆化搜索。

在使用记忆化搜索的时候，其实可以发现，区间中的元素是有序并连续的。相同长度的区间构成的BST的数量是相同的。
例如：[1,2,3]与[3,4,5]构成的BST数量相同。因此可以用区间的长度作为mem数组的key

下面的题解写的不错，使用了动态规划
https://leetcode.cn/problems/unique-binary-search-trees/solutions/6693/hua-jie-suan-fa-96-bu-tong-de-er-cha-sou-suo-shu-b/


递归解法
*/

func numTrees(n int) int {

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

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees(3))
}

/*
动态规划解法
*/

func numTreesDp(n int) int {

	return 0

}
