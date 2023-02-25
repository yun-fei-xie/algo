package _7_binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {

	var insert func(node *TreeNode, val int) *TreeNode
	insert = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		}
		if node.Val > val {
			node.Left = insert(node.Left, val)

		} else if node.Val < val {
			node.Right = insert(node.Right, val)
		}
		return node
	}
	root = insert(root, val)
	return root
}
