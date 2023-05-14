package binaryTree

import "container/list"

/*
中序遍历
*/
func inorderTraversal(root *treeNode) []int {

	values := make([]int, 0)
	stack := list.New()
	if root == nil {
		return values
	}

	cur := root

	for stack.Len() > 0 || cur != nil {

		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			node := stack.Remove(stack.Back()).(*treeNode)
			values = append(values, node.Val)
			cur = node.Right
		}

	}
	return values
}
