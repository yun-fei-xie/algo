package treeDP

import "math"

/*
网宿科技3月7号笔试题

https://leetcode.cn/problems/binary-tree-maximum-path-sum/?favorite=2cktkvj

路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。


方法：以当前节点为根节点的最大路径等于（左子树的最大路径和+右子树的最大路径和）


1. 计算当前内部路径最大值
2. 向上输出，只能输出当前节点+left 与当前节点+right的最大值

参考题解：https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/297276/shou-hui-tu-jie-hen-you-ya-de-yi-dao-dfsti-by-hyj8/

*/

func maxPathSum1(root *TreeNode) int {
	var ans = math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		ans = max(ans, left+right+node.Val)

		return max(max(left, right)+node.Val, 0)
	}
	dfs(root)
	return ans

}

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}
