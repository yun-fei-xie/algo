package _7_binarytree

/*
https://leetcode.cn/problems/sum-of-left-leaves/

给定二叉树的根节点，返回所有的左叶子之和

1. 叶子节点(该节点的左右孩子为nil)
2. 它是父节点的左孩子(这一点需要从父节点进行判断)

** 无非就是找出符合条件的节点 **
那么就用最简单的遍历方式，对每一个遇到的节点进行判断。

只有一个节点，不能算左叶子节点
*/
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}
	res := 0

	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 左叶子的定义
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			res += node.Left.Val
		}

		preOrder(node.Left)
		preOrder(node.Right)

	}

	preOrder(root)

	return res
}
