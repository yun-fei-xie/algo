package _7_binarytree

/*
https://leetcode.cn/problems/binary-tree-postorder-traversal/

*/

func postorderTraversal(root *TreeNode) []int {

	values := make([]int, 0)
	var postorder func(node *TreeNode)

	postorder = func(node *TreeNode) {

		if node == nil {
			return
		}

		postorder(node.Left)
		postorder(node.Right)
		values = append(values, node.Val)
	}

	postorder(root)
	return values
}
