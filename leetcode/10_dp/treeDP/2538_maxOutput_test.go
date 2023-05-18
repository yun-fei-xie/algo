package treeDP

import (
	"fmt"
	"testing"
)

/*
2538. 最大价值和与最小价值和的差值
https://leetcode.cn/problems/difference-between-maximum-and-minimum-price-sum/
给你一个 n 个节点的无向无根图，节点编号为 0 到 n - 1 。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间有一条边。
每个节点都有一个价值。给你一个整数数组 price ，其中 price[i] 是第 i 个节点的价值。
一条路径的 价值和 是这条路径上所有节点的价值之和。
你可以选择树中任意一个节点作为根节点 root 。选择 root 为根的 开销 是以 root 为起点的所有路径中，价值和 最大的一条路径与最小的一条路径的差值。
请你返回所有节点作为根节点的选择中，最大 的 开销 为多少。



*/

func maxOutput(n int, edges [][]int, price []int) int64 {
	var tree = make([][]int, n)
	for i := 0; i < len(edges); i++ {
		par, child := edges[i][0], edges[i][1]
		tree[par] = append(tree[par], child)
	}
	fmt.Println(tree)
	return -1
}

func TestMaxOutput(t *testing.T) {
	fmt.Println(maxOutput(6, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}}, []int{9, 8, 7, 6, 10, 5}))
}
