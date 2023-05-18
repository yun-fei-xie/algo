package treeDP

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
687. 最长同值路径
https://leetcode.cn/problems/longest-univalue-path/description/
给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
两个节点之间的路径长度 由它们之间的边数表示。

方法：


*/

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 0
		}
		var m, retm, left, right int
		if node.Left != nil {
			left = dfs(node.Left)
			if node.Val == node.Left.Val {
				m += left + 1
				retm = max(retm, left+1)
			}
		}
		if node.Right != nil {
			right = dfs(node.Right)
			if node.Val == node.Right.Val {
				m += right + 1
				retm = max(retm, right+1)
			}
		}
		// 所有子树中的最大值就是最后的答案
		ans = max(ans, m)
		return retm
	}
	dfs(root)
	return ans
}
