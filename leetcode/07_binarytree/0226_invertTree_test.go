package _7_binarytree

/*
https://leetcode.cn/problems/invert-binary-tree/description/

翻转二叉树

迭代法怎么做？
*/

func invertTree(root *TreeNode) *TreeNode {
	return invertAux(root)
}

func invertAux(node *TreeNode) *TreeNode {

	if node == nil {
		return nil
	}

	left := invertTree(node.Left)
	right := invertTree(node.Right)

	node.Left, node.Right = right, left
	return node
}
