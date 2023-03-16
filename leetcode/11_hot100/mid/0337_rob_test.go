package mid

import "math"

/*

https://leetcode.cn/problems/house-robber-iii/?favorite=2cktkvj

小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。


1.尝试暴力解法,从根节点开始，可能偷也可以不偷，分类枚举



*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
这个模型明显是一个后续，需要在根节点对数据进行汇总
*/

func rob3(root *TreeNode) int {

	var search func(node *TreeNode) int
	search = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		robIncludeNode := node.Val
		if node.Left != nil {
			robIncludeNode += search(node.Left.Left) + search(node.Left.Right)
		}
		if node.Right != nil {
			robIncludeNode += search(node.Right.Left) + search(node.Right.Right)
		}

		robExcludeNode := rob(node.Left) + rob(node.Right)

		if robIncludeNode > robExcludeNode {
			return robIncludeNode
		} else {
			return robExcludeNode
		}
	}

	return search(root)

}

// 加入记忆化搜索 有2个测试超时

func rob4(root *TreeNode) int {

	mem := make(map[*TreeNode]int) // 记录以key为根节点，打劫可以获得的最大收益

	var search func(node *TreeNode) int
	search = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		if _, found := mem[node]; found { //
			return mem[node]
		}

		robIncludeNode := node.Val
		if node.Left != nil {
			robIncludeNode += search(node.Left.Left) + search(node.Left.Right)
		}
		if node.Right != nil {
			robIncludeNode += search(node.Right.Left) + search(node.Right.Right)
		}

		robExcludeNode := rob(node.Left) + rob(node.Right)

		res := 0
		if robIncludeNode > robExcludeNode {
			res = robIncludeNode
		} else {
			res = robExcludeNode
		}
		mem[node] = res
		return res
	}

	return search(root)

}

/*
树形动态规划,不像常规dp在迭代中填表格，而是在递归遍历中填写表格
打劫一个树的最大收益，是 robIncludeRoot 和 robExcludeRoot 中的较大者。
即每个子树都有两个状态下的最优解：没打劫 root、和有打劫 root 下的最优解。

有两个变量共同决定一个状态：1、代表不同子树的 root 节点、2、是否打劫了 root。
可以维护一个二维数组 dp，但对象不能作为数组索引，改用 map。key 是当前子树的 root 节点，value 是存放两个状态的 res 数组。

没打劫根节点，则左右子树的根节点可打劫可不打劫： res[0] = 左子树的两个状态的较大值 + 右子树的两个状态的较大值。
打劫了根节点，则左右子树的根节点不能打劫： res[1] = root.val + 左子树的 [0] 状态 + 右子树的 [0] 状态。

*/

func rob3dp(root *TreeNode) int {

	dp := make(map[*TreeNode][]int)
	var max func(i int, j int) int
	max = func(i int, j int) int {
		if i > j {
			return i
		}
		return j
	}
	var search func(node *TreeNode) []int
	search = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		// 打劫左右子树
		left := search(node.Left)
		right := search(node.Right)

		// 之前没有遍历过当前节点
		if _, found := dp[node]; !found {
			dp[node] = []int{0, 0}
		}

		// 获取当前节点的res数组
		res, _ := dp[node]
		res[0] = max(left[0], left[1]) + max(right[0], right[1])
		res[1] = node.Val + left[0] + right[0]

		return res
	}

	res := search(root)
	return max(res[0], res[1])

}

func rob(root *TreeNode) int {
	// 思路乱了 这里的代码是错误的

	var search func(node *TreeNode, isRob bool) int
	search = func(node *TreeNode, isRob bool) int {
		if node == nil {
			return 0
		}
		if isRob { // 当前节点可以偷->并且偷当前节点

			left := search(node.Left, false)
			right := search(node.Right, false)
			return left + right + node.Val
		} else { // 下一轮可偷可不偷 会出现4种可能 -> 取最大
			left := search(node.Left, true)
			right := search(node.Right, true)
			return left + right
		}
	}

	return int(math.Max(float64(search(root, false)), float64(search(root, true))))
}
