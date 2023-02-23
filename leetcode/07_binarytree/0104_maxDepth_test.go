package _7_binarytree

/*
https://leetcode.cn/problems/maximum-depth-of-binary-tree/

给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

此题有多种解法，除了递归（代码量极少）
还可以用层序遍历的方法（遍历多少层 最大深度就是多少）

*/

func maxDepth(root *TreeNode) int {
	return maxDepthAux(root)
}

func maxDepthAux(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := maxDepthAux(node.Left)
	right := maxDepthAux(node.Right)

	return max(left, right) + 1

}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j

}
