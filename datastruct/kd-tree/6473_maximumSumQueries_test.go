package _4_kd_tree_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

/*
6473.最大和查询
https://leetcode.cn/problems/maximum-sum-queries/
给你两个长度为 n 、下标从 0 开始的整数数组 nums1 和 nums2 ，
另给你一个下标从 1 开始的二维数组 queries ，其中 queries[i] = [xi, yi] 。
对于第 i 个查询，在所有满足 nums1[j] >= xi 且 nums2[j] >= yi 的下标 j (0 <= j < n) 中，
找出 nums1[j] + nums2[j] 的 最大值 ，如果不存在满足条件的 j 则返回 -1 。
返回数组 answer ，其中 answer[i] 是第 i 个查询的答案。

思路：
0.题意比较容易理解，对于询问q{k,r} 需要nums1[j]>=k && nums2[j]>=r的提前下，求最大值。
1.最后的答案和j没有关系(不是求j的最大值，比赛的时候理解错误题意。而是求nums1[j]+nums2[j]的最大值)，也就是和元素的顺序没有关系。nums1和nums2中的每一个元素的位置一一对应（类似二维平面的点）

方法：离线查询问题
2-d树实现 暴力插入容易导致树的高度变高。
基于partition构建树也超时
1. 按照轴进行排序 然后选择中间进行递归
2. 当区间中只有一个元素的时候停止递归

*/

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	//构造点集
	length := len(nums1)
	points := make([]*point, length)
	for i := 0; i < length; i++ {
		points[i] = &point{x: nums1[i], y: nums2[i], sum: nums1[i] + nums2[i]}
	}

	// 插入2-d树
	tree := NewKdTree(points)
	var ans = []int{}
	// range查询
	for _, query := range queries {
		maxSum := tree.RangeQuery(&RectHV{
			xmin: query[0],
			ymin: query[1],
			xmax: math.MaxInt,
			ymax: math.MaxInt,
		})
		ans = append(ans, maxSum)
	}
	return ans
}

// 2-d树节点对应的矩形空间
type RectHV struct {
	xmin, ymin, xmax, ymax int
}

// 二维平面的一个坐标点 sum = x + y
type point struct {
	x   int
	y   int
	sum int
}

// 1.2-d树的节点
type kdNode struct {
	point *point
	lb    *kdNode
	rt    *kdNode
}

type kdTree struct {
	root *kdNode
}

func NewKdTree(points []*point) *kdTree {
	kt := &kdTree{root: nil}
	kt.root = kt.insert(points, true)
	return kt
}

func (kt *kdTree) insert(pts []*point, evenLevel bool) *kdNode {
	if len(pts) <= 1 {
		return &kdNode{
			point: pts[0],
		}
	}
	// 如果是x轴 按照x排序 否则按照y排序
	mid := (len(pts) - 1) / 2
	pivotKthLargest(pts, evenLevel, mid)

	var leftTree, rightTree *kdNode
	if len(pts[:mid]) >= 1 {
		leftTree = kt.insert(pts[:mid+1], !evenLevel) // [0...mid) 右边是开区间
	}
	if len(pts[mid+1:]) >= 1 {
		rightTree = kt.insert(pts[mid+1:], !evenLevel)
	}
	return &kdNode{
		point: pts[mid],
		lb:    leftTree,
		rt:    rightTree,
	}
}

// 2-D tree本质是二分搜索，相对于普通的BST增加了一个维度。
// 所以，思想还是和一维的BST相似。
// 从根节点出发，首先检查当前节点node是否在搜索区域中？如果是，则收入囊中。
// 然后检查切分轴与RectHV中对应轴的位置关系：
// 1. 切分轴在RectHV对应轴的左侧，则去node的右子树中递归查找（如果右子树不为空）
// 2. 切分轴在RectHV对应轴的右侧，则去node的左子树中查找（如果左子树不为空）
// 3. 切分轴在RectHV对应轴范围的中间。则需要去左右子树中递归查找。
func (kt *kdTree) RangeQuery(hv *RectHV) int {
	var ans = -1 // -1代表没查到的值
	var query func(node *kdNode, queryRange *RectHV, evenLevel bool)
	query = func(node *kdNode, queryRange *RectHV, evenLevel bool) {
		if node == nil {
			return
		}
		pt := node.point
		if pt.x >= queryRange.xmin && pt.x <= queryRange.xmax && pt.y >= queryRange.ymin && pt.y <= queryRange.ymax {
			ans = max(ans, pt.sum)
		}
		// 如果是x轴 切分轴在查询区间左侧 应该去右子树查
		if evenLevel && node.point.x <= queryRange.xmin {
			query(node.rt, queryRange, !evenLevel)
			return
		}
		// 如果是x轴 切分轴在查询区间右侧 应该去左子树查
		if evenLevel && node.point.x >= queryRange.xmax {
			query(node.lb, queryRange, !evenLevel)
			return
		}
		// 如果是x轴 切分轴在查询区间中间 递归去左右子树查
		if evenLevel && node.point.x < queryRange.xmax && node.point.x > queryRange.xmin {
			query(node.rt, queryRange, !evenLevel)
			query(node.lb, queryRange, !evenLevel)
			return
		}

		// 对于y轴同理
		if !evenLevel && node.point.y <= queryRange.ymin {
			query(node.rt, queryRange, !evenLevel)
			return
		}
		if !evenLevel && node.point.y >= queryRange.ymax {
			query(node.lb, queryRange, !evenLevel)
			return
		}
		if !evenLevel && node.point.y < queryRange.ymax && node.point.y > queryRange.ymin {
			query(node.rt, queryRange, !evenLevel)
			query(node.lb, queryRange, !evenLevel)
			return
		}

	}
	query(kt.root, hv, true)
	return ans
}

func pivotKthLargest(pts []*point, evenLevel bool, k int) {
	length := len(pts)
	left := 0
	right := length - 1
	for {
		pivotIndex := partition(pts, left, right, evenLevel)
		if pivotIndex == k {
			return
		} else if pivotIndex > k {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

/*
[left , right] 闭区间
*/
func partition(pts []*point, left int, right int, evenLevel bool) int {
	//pivotNum := arr[left]
	//随机 让索引落到[left,right]
	randIndex := rand.Intn(right-left+1) + left
	pts[randIndex], pts[left] = pts[left], pts[randIndex]
	var pivotNum int
	if evenLevel {
		pivotNum = pts[left].x
	} else {
		pivotNum = pts[left].y
	}
	i := left + 1
	j := right
	for i <= j {
		if (evenLevel && pts[j].x > pivotNum) || (!evenLevel && pts[j].y > pivotNum) {
			j--
		} else if (evenLevel && pts[j].x <= pivotNum) || (!evenLevel && pts[j].y <= pivotNum) {
			i++
		} else if (evenLevel && pts[j].x <= pivotNum) || (!evenLevel && pts[j].y <= pivotNum) {
			pts[j], pts[i] = pts[i], pts[j]
			i++
		} else if (evenLevel && pts[i].x > pivotNum) || (!evenLevel && pts[i].y > pivotNum) {
			pts[j], pts[i] = pts[i], pts[j]
			j--
		}
	}
	pts[i-1], pts[left] = pts[left], pts[i-1]
	return i - 1
}

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

func TestMaximumSumQueries(t *testing.T) {
	//fmt.Println(maximumSumQueries([]int{3, 2, 5}, []int{2, 3, 4}, [][]int{{4, 4}, {3, 2}, {1, 1}}))
	//fmt.Println(maximumSumQueries([]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 5}}))
	//fmt.Println(maximumSumQueries([]int{2, 1}, []int{2, 3}, [][]int{{3, 3}}))
	//fmt.Println(maximumSumQueries([]int{5, 11, 20, 80, 95, 14, 44, 26, 21, 6, 94, 33, 40, 2, 94, 89},
	//	[]int{60, 76, 61, 6, 7, 71, 22, 26, 100, 63, 17, 2, 89, 19, 100, 69}, [][]int{{80, 18}, {99, 27}, {11, 16}, {36, 86}, {98, 80}, {83, 15}, {47, 31}, {7, 70}, {94, 64}, {16, 48}, {33, 41}, {89, 86}, {61, 54}, {25, 40}, {50, 9}, {38, 84}, {30, 77}, {78, 19}, {88, 15}}))
	fmt.Println(maximumSumQueries([]int{93, 53, 5, 6, 12}, []int{48, 87, 39, 72, 90}, [][]int{{69, 75}, {22, 39}, {53, 81}, {90, 48}}))
}
