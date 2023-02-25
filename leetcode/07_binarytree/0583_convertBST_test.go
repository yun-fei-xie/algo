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
https://leetcode.cn/problems/convert-bst-to-greater-tree/description/
反中序遍历(右子树 当前节点 左子树)，倒着累加
*/
func convertBST(root *TreeNode) *TreeNode {

	var pre *TreeNode

	var convertBSTAux func(node *TreeNode)
	convertBSTAux = func(node *TreeNode) {
		if node == nil {
			return
		}

		convertBSTAux(node.Right)
		if pre != nil {
			node.Val += pre.Val
		}
		pre = node
		convertBSTAux(node.Left)
	}
	convertBSTAux(root)

	return root
}
