package _7_binarytree

/*
https://leetcode.cn/problems/symmetric-tree/

对比一颗树的左右子树是否对称
这题体会不够深
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compareTree(root.Left, root.Right)

}

func compareTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}

	// left.val == right.val

	boolLeft := compareTree(left.Left, right.Right)
	boolRight := compareTree(left.Right, right.Left)

	return boolLeft && boolRight
}
