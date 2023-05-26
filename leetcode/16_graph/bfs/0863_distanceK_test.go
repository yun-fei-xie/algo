package bfs_test

import (
	"container/list"
)

/*
863. 二叉树中所有距离为 K 的结点
https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/description/?envType=study-plan-v2&envId=graph-theory

给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。
返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。

方法：把二叉树转换成图，然后用图的广度优先遍历。(题目说了，节点值各不相同 节点范围1-500)
1.如何将二叉树转换成图？在遍历的过程中，建立关系
2.不遍历，不知道节点有多少？节点值在[0...500]也就是最多501个节点

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	graph := make([][]int, 501)
	for i := 0; i < 501; i++ {
		graph[i] = make([]int, 0)
	}

	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
		}
		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)

	// bfs
	visited := make([]bool, 501)
	queue := list.New()
	queue.PushBack(target.Val)
	visited[target.Val] = true
	step := 0
	ans := make([]int, 0)
	for queue.Len() != 0 {

		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {
			vertex := queue.Remove(queue.Front()).(int)
			if step == k {
				ans = append(ans, vertex)
			}
			for _, adj := range graph[vertex] {
				if !visited[adj] {
					visited[adj] = true
					queue.PushBack(adj)
				}
			}
		}
		if step == k {
			break
		}

		step++
	}
	return ans
}
