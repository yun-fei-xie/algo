package _7_binarytree_test

import "container/list"

/*
https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

在层序遍历的时候，如果一个节点的左右节点都是空，那么就找到了一个最小深度
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	minDep := 0

	for queue.Len() > 0 {

		levelCount := queue.Len()
		minDep++
		for i := 0; i < levelCount; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return minDep
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}
	}
	return minDep
}
