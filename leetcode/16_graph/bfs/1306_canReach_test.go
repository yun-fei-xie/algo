package bfs_test

import "container/list"

/*
1306. 跳跃游戏 III
https://leetcode.cn/problems/jump-game-iii/description/

这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。
请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。
注意，不管是什么情况下，你都无法跳到数组之外。

方法：从一个位置i可以跳向另外一个位置j可以表示为从i到j存在一条有向边（i->j）
基于这个方法可以将数组的每一个下标抽象成图中的一个一个的顶点。
跳跃的可达性可以表示为图中顶点的边。
这样的话，就可以使用广度优先遍历，把图中的顶点遍历一遍。如果碰到arr[vertex]==0 就表示可以调到对应元素值为0的下标处。
*/
func canReach(arr []int, start int) bool {
	// 每一个位置都有一个可达性（可能是一个顶点，也可能是两个订点）
	// 那么可以相应地构建一副有向图
	arrLen := len(arr)
	graph := make([][]int, arrLen)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < arrLen; i++ {
		if i+arr[i] >= 0 && i+arr[i] < arrLen {
			graph[i] = append(graph[i], i+arr[i])
		}
		if i-arr[i] >= 0 && i-arr[i] < arrLen {
			graph[i] = append(graph[i], i-arr[i])
		}
	}

	// bfs
	visited := make([]bool, arrLen)
	visited[start] = true
	queue := list.New()
	queue.PushBack(start)
	for queue.Len() != 0 {

		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {

			vertex := queue.Remove(queue.Front()).(int)
			if arr[vertex] == 0 {
				return true
			}

			for _, adj := range graph[vertex] {
				if !visited[adj] {
					queue.PushBack(adj)
					visited[adj] = true
				}
			}
		}
	}

	return false
}
