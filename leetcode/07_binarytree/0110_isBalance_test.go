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
https://leetcode.cn/problems/balanced-binary-tree/description/

本质是求高度，不过感觉代码有些冗余
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, res := isBalanceAux(root)
	return res
}

func isBalanceAux(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftIsBalance := isBalanceAux(node.Left)
	rightHeight, rightIsBalance := isBalanceAux(node.Right)

	if abs(leftHeight, rightHeight) <= 1 && leftIsBalance && rightIsBalance {
		return maxNumber(leftHeight, rightHeight) + 1, true
	} else {
		return maxNumber(leftHeight, rightHeight) + 1, false
	}
}

func maxNumber(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}

}

func abs(i, j int) int {
	if i > j {
		return i - j
	} else {
		return j - i
	}
}
