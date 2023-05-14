package _7_binarytree

/*
https://leetcode.cn/problems/same-tree/

给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

方法1：
如何定义两棵树相同。
1. t1和t2的根节点相同
2. t1的左子树=t2的左子树  t1的右子树=t2的右子树

递归逻辑：于是可以将原问题拆解为子问题。（先看根节点是不是相同，如果相同看看左右子树是否相同）
基本case:两个节点同时为空
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil || p != nil && q == nil {
		return false
	} else if p == nil && q == nil {
		return true
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
