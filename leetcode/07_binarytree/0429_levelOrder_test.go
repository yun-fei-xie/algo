package _7_binarytree

/*
https://leetcode.cn/problems/n-ary-tree-level-order-traversal/

N叉树的层序遍历
*/
//type Node struct {
//	Val      int
//	Children []*Node
//}
//
//func levelOrderNtree(root *Node) [][]int {
//
//	values := make([][]int, 0)
//
//	var levelorder func(node *Node, depth int)
//	levelorder = func(node *Node, depth int) {
//
//		if node == nil {
//			return
//		}
//		if depth == len(values) {
//			values = append(values, []int{})
//		}
//
//		values[depth] = append(values[depth], node.Val)
//		for i := 0; i < len(node.Children); i++ {
//			child := node.Children[i]
//			levelorder(child, depth+1)
//		}
//
//	}
//
//	levelorder(root, 0)
//	return values
//
//}
