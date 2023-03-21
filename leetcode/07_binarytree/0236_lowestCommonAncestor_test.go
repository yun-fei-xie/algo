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
https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

二叉树最近公共祖先问题
这个题的精髓在于后续遍历过程中返回哪个节点？
还有一点是，当递归到p节点或者q节点时，直接返回，不需要继续向下递归

递归何时返回-> 1.节点为空、遇到了节点p、遇到了节点q
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var postOrder func(node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode
	postOrder = func(node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
		if node == nil || node == p || node == q {
			return node
		}

		left := postOrder(node.Left, p, q)
		right := postOrder(node.Right, p, q)

		// 三种情况
		if left != nil && right != nil { // 1.左右都不为空，自底向上看的话，当前节点就是最小公共祖先
			return node
		} else if left != nil {
			return left
		} else if right != nil {
			return right
		}
		return nil
	}
	return postOrder(root, p, q)
}
