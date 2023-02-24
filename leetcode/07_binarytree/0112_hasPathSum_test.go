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

路径之和与字符串那个（257号 二叉树的路径）十分接近


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
