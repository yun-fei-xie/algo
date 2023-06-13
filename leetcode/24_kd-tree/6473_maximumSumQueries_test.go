package _4_kd_tree_test

import (
	"fmt"
	"math"
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

func (rt *RectHV) Contains(other *RectHV) bool {
	if rt.xmin <= other.xmin && rt.xmax >= rt.xmax && rt.ymin <= other.ymin && rt.ymax >= other.ymax {
		return true
	}
	return false
}

// 二维平面的一个坐标点 sum = x + y
type point struct {
	x   int
	y   int
	sum int
}

func (p *point) equals(other *point) bool {
	if p.x == other.x && p.y == other.y {
		return true
	}
	return false
}

// 1.2-d树的节点
type kdNode struct {
	point  *point
	rectHv *RectHV
	lb     *kdNode
	rt     *kdNode
}

type kdTree struct {
	root *kdNode
	size int
}

// 将点集合构造成一颗2d-tree 初始时，每个点管理的矩形是整个二维平面
// 这个二维平面在节点插入的过程中不断被修正
func NewKdTree(points []*point) *kdTree {
	size := len(points)
	kt := &kdTree{root: nil, size: size}
	for i := 0; i < size; i++ {
		kt.root = kt.insert(kt.root, points[i], true, &RectHV{
			xmin: math.MinInt,
			ymin: math.MinInt,
			xmax: math.MaxInt,
			ymax: math.MaxInt,
		})
	}
	return kt
}

func (kt *kdTree) insert(root *kdNode, pt *point, evenLevel bool, rectHv *RectHV) *kdNode {
	if root == nil {
		return &kdNode{point: pt, rectHv: rectHv}
	}
	pos := comparePoints(root, pt, evenLevel)
	// 如果root划分的是x轴，并且pt在root的左边。更新pt的管理区域x坐标的上界。rectHv.xmax=root.x
	if pos < 0 && evenLevel {
		rectHv.xmax = root.point.x
		root.lb = kt.insert(root.lb, pt, !evenLevel, rectHv)
	} else if pos > 0 && evenLevel {
		// 如果root划分的是x轴，并且pt在root的右边。更新pt管理区域x坐标的下界。rectHv.xmin=root.x
		rectHv.xmin = root.point.x
		root.rt = kt.insert(root.rt, pt, !evenLevel, rectHv)
	} else if pos < 0 && !evenLevel {
		// 如果root划分的是y轴，并且pt在root的下边。更新pt管理区域y坐标的上界。rectHv.ymax =root.y
		rectHv.ymax = root.point.y
		root.lb = kt.insert(root.lb, pt, !evenLevel, rectHv)
	} else if pos > 0 && !evenLevel {
		//如果root划分的是y轴，并且pt在root的上边。更新pt管理区域y坐标的下界。rectHv.ymin=root.y
		rectHv.ymin = root.point.y
		root.rt = kt.insert(root.rt, pt, !evenLevel, rectHv)
	}

	// 处理pos等于0的情况并且两个点不重合 统一安排到右子树。这个时候pt管理的矩形区域和root相同。（注意体会）
	if pos == 0 && !pt.equals(root.point) {
		root.rt = kt.insert(root.rt, pt, !evenLevel, rectHv)
	}
	return root
}

// 比较两个点的相对位置，如果是偶数层就比较x,如果是奇数层就比较y。
// 如果p比node小(ans <=0),p应该去ans的左孩子递归插入；如果p比node大（ans >0),p应该去ans的右孩子插入。
// 也就是说，如果node划分的是x轴（把空间划分成了左右两部分)那么比较p和node的x值，决定p应该去node管理的空间的左边还是右边。
// 如果node划分的是y轴（把空间划分成了上下两个部分）那么比较p和node的y值，决定p应该去node管理的空间的上边还是下边。
func comparePoints(node *kdNode, p *point, evenLevel bool) (ans int) {
	if evenLevel {
		return p.x - node.point.x
	} else {
		return p.y - node.point.y
	}
}

// 2-D tree本质是二分搜索，相对于普通的BST增加了一个维度。
// 所以，思想还是和一维的BST相似。
// 从根节点出发，首先检查当前节点node是否在搜索区域中？如果是，则收入囊中。
// 然后检查切分轴与RectHV中对应轴的位置关系：1. 切分轴在RectHV对应轴的左侧，则去node的右子树中递归查找（如果右子树不为空）
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
		// 如果是x轴 切分轴在查询区间左侧 应该去右子树查 (不需要更新查询区间)
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
