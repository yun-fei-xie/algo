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
https://leetcode.cn/problems/subtree-of-another-tree/description/

给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。
二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。

整个过程分为2步：
1. 对root进行遍历
2. 遍历过程中对每个节点进行一次子树对比（是否是相同的树）

*/
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var preOrder func(node *TreeNode)
	var result bool
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		comp := isSubTreeAux(node, subRoot)
		if comp == true {
			result = true
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)

	return result
}

func isSubTreeAux(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isSubTreeAux(left.Left, right.Left) && isSubTreeAux(left.Right, right.Right)
}
