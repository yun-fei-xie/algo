package _7_binarytree

import (
	"container/list"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
103. 二叉树的锯齿形层序遍历

给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

方法：层序遍历
碰到偶数层，将该层的数组进行翻转
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	// 这里需要做特殊判断 不然在queue中压入root(nil),再从队列中取出节点，执行node.left 或者是node.right会产生空指针异常
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	even := false
	for queue.Len() != 0 {
		l := queue.Len()
		level := make([]int, l)
		for i := 0; i < l; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			level[i] = node.Val
		}

		if even {
			reverseSlice(level)
			even = false
		} else {
			even = true
		}
		ans = append(ans, level)
	}
	return ans

}

func reverseSlice(numbers []int) {
	left, right := 0, len(numbers)-1
	for left <= right {
		numbers[left], numbers[right] = numbers[right], numbers[left]
		left++
		right--
	}
}
