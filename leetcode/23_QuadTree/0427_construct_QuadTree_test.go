package _3_QuadTree

import (
	"fmt"
	"testing"
)

/*
427. 建立四叉树
https://leetcode.cn/problems/construct-quad-tree/description/

给你一个 n * n 矩阵 grid ，矩阵由若干 0 和 1 组成。请你用四叉树表示该矩阵 grid 。
你需要返回能表示矩阵的 四叉树 的根结点。
注意，当 isLeaf 为 False 时，你可以把 True 或者 False 赋值给节点，两种值都会被判题机制 接受 。
四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：

val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False；
isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。

方法：递归构建即可
*/

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {

	row, col := len(grid), len(grid[0])
	// 判断一个区域是否已经构成叶子节点
	var isLeaf func(rowL int, rowR int, colL int, colR int) bool
	isLeaf = func(rowL int, rowR int, colL int, colR int) bool {
		start := grid[rowL][colL]
		for i := rowL; i <= rowR; i++ {
			for j := colL; j <= colR; j++ {
				if grid[i][j] != start {
					return false
				}
			}
		}
		return true
	}
	// 传入一个二维平面的范围，返回这个范围内的四叉树的根节点
	var helper func(rowL int, rowR int, colL int, colR int) *Node
	helper = func(rowL int, rowR int, colL int, colR int) *Node {
		if isLeaf(rowL, rowR, colL, colR) {
			node := &Node{
				Val:         true,
				IsLeaf:      true,
				TopLeft:     nil,
				TopRight:    nil,
				BottomLeft:  nil,
				BottomRight: nil,
			}
			if grid[rowL][colL] == 0 {
				node.Val = false
			}
			return node
		}

		//	将当前矩形递归划分成4分
		rowMid := rowL + (rowR-rowL)/2
		colMid := colL + (colR-colL)/2

		root := &Node{
			Val:         false,
			IsLeaf:      false,
			TopLeft:     helper(rowL, rowMid, colL, colMid),
			TopRight:    helper(rowL, rowMid, colMid+1, colR),
			BottomLeft:  helper(rowMid+1, rowR, colL, colMid),
			BottomRight: helper(rowMid+1, rowR, colMid+1, colR),
		}
		return root
	}

	return helper(0, row-1, 0, col-1)
}

func TestConstructor(t *testing.T) {
	fmt.Println(construct([][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0}}))
}
