package _4_kd_tree

import (
	"fmt"
	"sort"
	"testing"
)

/*
1828. 统计一个圆中点的数目
https://leetcode.cn/problems/queries-on-number-of-points-inside-a-circle/description/

方法：二维k-d树
*/

func countPoints(points [][]int, queries [][]int) []int {

	var cnt int
	var queryCircle func(root *node, circle []int, evenLevel bool)
	queryCircle = func(root *node, circle []int, evenLevel bool) {
		if root == nil {
			return
		}
		// 查看当前节点是否包含在圆中
		if contains(circle, root.point) {
			cnt++
		}
		// x轴
		if evenLevel && circle[0]+circle[2] < root.point[0] {
			queryCircle(root.left, circle, !evenLevel)
			return
		}
		if evenLevel && circle[0]-circle[2] > root.point[0] {
			queryCircle(root.right, circle, !evenLevel)
			return
		}
		if evenLevel && root.point[0] >= circle[0]-circle[2] && root.point[0] <= circle[0]+circle[2] {
			queryCircle(root.left, circle, !evenLevel)
			queryCircle(root.right, circle, !evenLevel)
			return
		}

		// y轴

		if !evenLevel && circle[1]+circle[2] < root.point[1] {
			queryCircle(root.left, circle, !evenLevel)
			return
		}
		if !evenLevel && circle[1]-circle[2] > root.point[1] {
			queryCircle(root.right, circle, !evenLevel)
			return
		}
		if !evenLevel && root.point[1] >= circle[1]-circle[2] && root.point[1] <= circle[1]+circle[2] {
			queryCircle(root.left, circle, !evenLevel)
			queryCircle(root.right, circle, !evenLevel)
			return
		}
	}
	kt := NewKdTree(points)
	var ans = make([]int, len(queries))
	for index, query := range queries {
		queryCircle(kt.root, query, true)
		ans[index] = cnt
		cnt = 0
	}
	return ans
}

type node struct {
	point []int //point[0]->x point[1]->y
	left  *node
	right *node
}

type kdTree struct {
	root *node
}

func NewKdTree(pts [][]int) *kdTree {

	var insert func(ps [][]int, evenLevel bool) *node
	insert = func(ps [][]int, evenLevel bool) *node {
		if len(ps) == 1 {
			return &node{point: ps[0]}
		}
		mid := (len(ps) - 1) / 2
		sort.Slice(ps, func(i, j int) bool {
			if evenLevel {
				return ps[i][0] < ps[j][0]
			} else {
				return ps[i][1] < ps[j][1]
			}
		})
		var leftTree, rightTree *node
		if len(ps[:mid]) >= 1 {
			leftTree = insert(ps[:mid], !evenLevel)
		}
		if len(ps[mid+1:]) >= 1 {
			rightTree = insert(ps[mid+1:], !evenLevel)
		}

		return &node{
			point: ps[mid],
			left:  leftTree,
			right: rightTree,
		}
	}

	root := insert(pts, true)
	return &kdTree{root: root}
}

func contains(circle []int, p []int) bool {
	d := (p[0]-circle[0])*(p[0]-circle[0]) + (p[1]-circle[1])*(p[1]-circle[1])
	if d <= circle[2]*circle[2] {
		return true
	}
	return false
}

func TestCountPoints(t *testing.T) {
	fmt.Println(countPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}))
}
