package _7_binarytree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
层序遍历-> 我怎么知道自己在最后一层？ 每一层迭代都记住最左边的数值。这个数值一直更新即可
*/
func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	res := 0
	for queue.Len() > 0 {
		len := queue.Len() // 这个地方是一个容易写错的地方
		for i := 0; i < len; i++ {

			node := queue.Remove(queue.Front()).(*TreeNode)

			if i == 0 {
				res = node.Val
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return res

}
