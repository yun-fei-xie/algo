package _7_binarytree

import (
	"container/list"
	"math"
)

/*
https://leetcode.cn/problems/find-largest-value-in-each-tree-row/

给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。


解法: 层序遍历中，内层循环每一轮会遍历一层，在这一轮中拿到最大值

*/

func largestValues(root *TreeNode) []int {

	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()
	levelCount := 1
	queue.PushBack(root)

	for queue.Len() > 0 {

		max := math.MinInt64
		nextLevelCount := 0

		for i := 0; i < levelCount; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
				nextLevelCount++
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
				nextLevelCount++
			}
		}
		levelCount = nextLevelCount
		res = append(res, max)
	}
	return res
}
