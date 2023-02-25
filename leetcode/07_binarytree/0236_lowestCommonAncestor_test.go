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
二叉树最近公共祖先问题
这个题的精髓在于后续遍历过程中返回哪个节点？
还有一点是，当递归到p节点或者q节点时，直接返回，不需要继续向下递归
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
		if left != nil && right != nil {
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
