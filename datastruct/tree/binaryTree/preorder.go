package binaryTree

import (
	"container/list"
	"fmt"
)

// 递归前序遍历
func preOrder(root *treeNode) {

	if root == nil {
		return
	}

	fmt.Println(root.Val)

	preOrder(root.Left)
	preOrder(root.Right)
}

// 非递归前序遍历
func preOrder2(root *treeNode) {
	stack := list.New()
	if root == nil {
		return
	}
	stack.PushBack(root)

	for stack.Len() != 0 {
		element := stack.Back()
		stack.Remove(element)
		node := element.Value.(*treeNode)
		fmt.Println(node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
}
