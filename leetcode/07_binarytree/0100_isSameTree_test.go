package _7_binarytree

/*
https://leetcode.cn/problems/same-tree/

给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeAux(p, q)
}

func isSameTreeAux(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isSameTreeAux(left.Left, right.Left) && isSameTreeAux(left.Right, right.Right)
}
