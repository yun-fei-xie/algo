package _7_binarytree

import "container/list"

/*
https://leetcode.cn/problems/count-complete-tree-nodes/description/

给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

如何不使用遍历，把时间复杂度压缩到O(1)
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	totalCounts := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {

		node := queue.Remove(queue.Front()).(*TreeNode)
		totalCounts++
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return totalCounts
}
