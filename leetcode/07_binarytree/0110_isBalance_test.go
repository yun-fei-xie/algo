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

判断一颗二叉树是否是平衡二叉树
1. 左子树是平衡的二叉树
2. 右子树是平衡的二叉树
3. 以当前节点为根的二叉树是平衡的二叉树

需要使用高度进行判断

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
