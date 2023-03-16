package hard

import "math"

/*
网宿科技3月7号笔试题

https://leetcode.cn/problems/binary-tree-maximum-path-sum/?favorite=2cktkvj

路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。


1. 计算当前内部路径最大值
2. 向上输出，只能输出当前节点+left 与当前节点+right的最大值

参考题解：https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/297276/shou-hui-tu-jie-hen-you-ya-de-yi-dao-dfsti-by-hyj8/

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var maxSum = math.MinInt32

	var max func(i, j int) int
	max = func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}

	}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		innerPathSum := node.Val + left + right
		maxSum = max(maxSum, innerPathSum)
		outPutPath := node.Val + max(left, right) // 向父节点进行输出的时候，当前子树的最大和可能为负数，负数只会减少上层的和
		return max(outPutPath, 0)

	}
	dfs(root)
	return maxSum
}
