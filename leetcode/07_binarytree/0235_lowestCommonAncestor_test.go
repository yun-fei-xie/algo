package _7_binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
/*
https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
二叉搜索树的最近公共祖先
如果当前节点的值比p、q都小，那么最近公共祖先在当前节点的右子树中
如果当前节点的值比p、q都大，那么最近公共祖先在当前节点的左子树中
如果当前节点的值在p、q之间，那么最近公共祖先就是当前节点。
*/
func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {

	var preOrder func(node, p, q *TreeNode) *TreeNode
	preOrder = func(node, p, q *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val > p.Val && node.Val > q.Val {
			return preOrder(node.Left, p, q)
		} else if node.Val < p.Val && node.Val < q.Val {
			return preOrder(node.Right, p, q)
		} else {
			return node
		}
	}

	return preOrder(root, p, q)
}
