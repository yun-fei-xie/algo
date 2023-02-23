package _7_binarytree

import (
	"container/list"
	"fmt"
	"testing"
)

/*
*
https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
用了闭包函数的写法 解决了values在递归函数中的共享问题
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {

	var values = make([]int, 0)

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {

		if node == nil {
			return
		}
		values = append(values, node.Val)
		preorder(node.Left)
		preorder(node.Right)

	}
	preorder(root)
	return values
}

func preorderTraversalIter(root *TreeNode) []int {

	values := make([]int, 0)
	stack := list.New()
	if root == nil {
		return values
	}
	stack.PushBack(root)

	for stack.Len() != 0 {
		element := stack.Back()
		stack.Remove(element)
		node := element.Value.(*TreeNode)
		values = append(values, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return values
}

func TestInorderTraversalIter(t *testing.T) {

	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	node1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node2 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	root.Right = node1
	node1.Left = node2

	res := preorderTraversalIter(root)
	fmt.Println(res)

}
