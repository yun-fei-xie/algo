package _7_binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
https://leetcode.cn/problems/path-sum/description/

给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
如果存在，返回 true ；否则，返回 false 。
叶子节点 是指没有子节点的节点。

解法：
使用前序递归遍历
路径之和与字符串那个（257号 二叉树的路径）十分接近  https://leetcode.cn/problems/binary-tree-paths/


*/
func hasPathSum(root *TreeNode, targetSum int) bool {

	// 前序遍历 从root到叶子
	var preorder func(node *TreeNode, sum int) bool
	preorder = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}
		if node.Left == nil && node.Right == nil {
			// 叶子
			if node.Val+sum == targetSum {
				return true
			} else {
				return false
			}
		}

		left := preorder(node.Left, sum+node.Val)
		right := preorder(node.Right, sum+node.Val)

		if left || right {
			return true
		} else {
			return false
		}
	}
	return preorder(root, 0)

}
