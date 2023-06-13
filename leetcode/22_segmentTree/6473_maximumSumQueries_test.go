package _2_segmentTree_test

import (
	"container/list"
	"fmt"
	"sort"
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
和之前笔试长亭科技的题目有点像
0. 朴素解法（直接暴力遍历 每次查询，时间复杂度O(n),q次查询 时间复杂度O(qn)）
1. 排序+单调栈+二分
2. 动态开点线段树
3. kd-tree
4. quad-tree
*/

// 朴素解法 对于每个查询 做一轮遍历。这样的解法可以得到 1403/1414的通过率
func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	var ans = make([]int, 0)
	length := len(nums1)
	for _, query := range queries {
		x := query[0]
		y := query[1]
		sum := -1
		for i := 0; i < length; i++ {
			if nums1[i] >= x && nums2[i] >= y {
				if nums1[i]+nums2[i] > sum {
					sum = nums1[i] + nums2[i]
				}
			}
		}
		ans = append(ans, sum)
	}
	return ans
}

// 排序，因为题目要求nums1[i] + nums2[i] 的整体和最大，并且要求nums1[i]>=x nums2[i]>=y
// 可以将nums1[i]和nums1[i]-> {nums1[i],nums2[i]}元素对进行排序。
// [4,3,1,2] [2,4,9,5] -> {{4,2},{3,4},{1,9},{2,5}} -> 按照nums1升序排序{{1,9}, {2,5}, {3,4},{4,2}}
// 然后使用二分查找，可以有效过滤一些不可能的元素对。 通过率1396 / 1414 (居然比朴素通过率还要低一些 可能是排序占用了额外的时间)
type p struct {
	n1  int
	n2  int
	sum int
}

func maximumSumQueries2(nums1 []int, nums2 []int, queries [][]int) []int {
	length := len(nums1)

	pair := make([]p, length)
	for i := 0; i < length; i++ {
		pair[i] = p{
			n1: nums1[i],
			n2: nums2[i],
		}
	}
	sort.Slice(pair, func(i, j int) bool {
		return pair[i].n1 < pair[j].n1
	})

	var binarySearch func(pair []p, x int) int // 找出p.n1>=x的第一个下标
	binarySearch = func(pair []p, x int) int {
		left, right := 0, len(pair)-1
		for left <= right {
			mid := left + (right-left)/2
			if pair[mid].n1 == x {
				return mid
			} else if pair[mid].n1 < x {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
		return left
	}
	var ans = make([]int, 0)
	for _, query := range queries {
		x := query[0]
		y := query[1]
		xIndex := binarySearch(pair, x)
		sum := -1
		if xIndex >= length {
			ans = append(ans, -1)
			continue
		} else {
			for j := xIndex; j < length; j++ {
				if pair[j].n2 >= y {
					if pair[j].n1+pair[j].n2 > sum {
						sum = pair[j].n1 + pair[j].n2
					}
				}
			}
		}
		ans = append(ans, sum)
	}
	return ans
}

// 在maximumSumQueries2的基础上，如果能够将（符合条件的）最大值固定在一个区间中，那么便可以使用一些区间查询的方式快速找到最大值。
// 在maximumSumQueries2的中，对第一维进行升序排序后，可以发现，n1>=x的值都在数组的右边，此时，如果能够在对n1排序的基础上对n2进行降序排序，
// 那么n2>=y的值都会卡在数组的左边。左右相夹，就可以得到一个区间[i...j]。[i...len-1]这一段都是满足>=x，[0...j]这一段都满足>=y
// 那么[i...j]这一段属于区间交集。既满足>=x 也满足 >=y 。也就是要求这个区间中的最大值。（可以用线段树或者是树状数组）
func maximumSumQueries3(nums1 []int, nums2 []int, queries [][]int) []int {
	length := len(nums1)
	pair := make([]*p, length)
	for i := 0; i < length; i++ {
		pair[i] = &p{
			n1:  nums1[i],
			n2:  nums2[i],
			sum: nums1[i] + nums2[i],
		}
	}
	sort.Slice(pair, func(i, j int) bool {
		if pair[i].n1 != pair[j].n1 {
			return pair[i].n1 <= pair[j].n1
		} else {
			return pair[i].n2 >= pair[j].n2
		}
	})

	// 删除部分元素使得pair.n2呈现降序 (这一步也颇为麻烦) 因为是从后向前进行 [{2,3},{3,2},{5,4}] 这一步需要用到栈进行收集
	// 因为从右到左，每次需要和当前最左边一个没被删除的元素进行比较，这个时候只有栈比较方便。（和栈顶元素比较n2，比栈顶大则入栈，否则丢弃）
	stack := list.New()
	stack.PushBack(length - 1) // 存放下标即可
	delMark := make([]bool, length)
	for j := length - 2; j >= 0; j-- {
		if pair[j].n2 > pair[stack.Back().Value.(int)].n2 {
			stack.PushBack(j)
		} else {
			delMark[j] = true
		}
	}
	// 用一个新的slice收集删除过后的pair
	newPair := make([]*p, 0, length)
	for i := 0; i < length; i++ {
		if !delMark[i] {
			newPair = append(newPair, pair[i])
		}
	}

	// newPair 真正做到n1和n2都有序
	// 考虑是否有重复数据
	var binarySearchX func(pair []*p, x int) int
	var binarySearchY func(pair []*p, y int) int

	// 查找pair.n1 >=x 的第一个下标 ,也就是在可能重复的数据中找pair.n1==x的第一个下标
	binarySearchX = func(pair []*p, x int) int {
		left, right := 0, len(pair)-1
		for left <= right {
			mid := left + (right-left)/2
			if pair[mid].n1 == x {
				// 向前看一眼
				if mid-1 >= 0 && pair[mid-1].n1 < x {
					return mid
				}
				right = mid - 1

			} else if pair[mid].n1 < x {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}
	binarySearchY = func(pair []*p, y int) int {
		left, right := 0, len(pair)-1
		for left <= right {
			mid := left + (right-left)/2
			if pair[mid].n2 == y {
				// 向右看一眼
				if mid+1 < len(pair) && pair[mid+1].n2 < y {
					return mid
				}
				left = mid + 1

			} else if pair[mid].n2 < y {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return right
	}

	segmentTree := buildTree(newPair)
	// 判断[l...r]是否越界
	var ans = make([]int, len(queries))
	for i := 0; i < len(queries); i++ {

		x := queries[i][0]
		y := queries[i][1]
		l := binarySearchX(newPair, x)
		r := binarySearchY(newPair, y)
		if l == len(newPair) || l < 0 || r == len(newPair) || r < 0 || l > r {
			ans[i] = -1
			continue
		}
		ans[i] = segmentTree.query(l, r)
	}
	return ans
}

// 线段树 不需要更新 只需要构建和查询
type segmentTree struct {
	tree []int
	size int
}

func buildTree(data []*p) *segmentTree {
	tree := make([]int, len(data)*4)

	var constructor func(treeIndex int, left int, right int)
	constructor = func(treeIndex int, left int, right int) {
		if left == right {
			tree[treeIndex] = data[left].sum
			return
		}
		leftChild := 2*treeIndex + 1
		rightChild := 2*treeIndex + 2
		mid := left + (right-left)/2

		constructor(leftChild, left, mid)
		constructor(rightChild, mid+1, right)
		tree[treeIndex] = max(tree[leftChild], tree[rightChild])
	}
	constructor(0, 0, len(data)-1)
	return &segmentTree{tree: tree, size: len(data)}
}

// 查询
func (st *segmentTree) query(left, right int) int {

	var queryRange func(treeIndex int, left int, right int, queryL int, queryR int) int
	queryRange = func(treeIndex int, left int, right int, queryL int, queryR int) int {
		if left == queryL && right == queryR {
			return st.tree[treeIndex]
		}
		mid := left + (right-left)/2
		leftChild := 2*treeIndex + 1
		rightChild := 2*treeIndex + 2

		if queryL > mid {
			return queryRange(rightChild, mid+1, right, queryL, queryR)
		} else if queryR <= mid {
			return queryRange(leftChild, left, mid, queryL, queryR)
		} else {
			return max(queryRange(leftChild, left, mid, queryL, mid), queryRange(rightChild, mid+1, right, mid+1, queryR))
		}
	}

	return queryRange(0, 0, st.size-1, left, right)
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
	//fmt.Println(maximumSumQueries2([]int{3, 2, 5}, []int{2, 3, 4}, [][]int{{4, 4}, {3, 2}, {1, 1}}))
	//fmt.Println(maximumSumQueries3([]int{3, 2, 5}, []int{2, 3, 4}, [][]int{{4, 4}, {3, 2}, {1, 1}}))
	//fmt.Println(maximumSumQueries3([]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 4}}))
	fmt.Println(maximumSumQueries2([]int{5, 11, 20, 80, 95, 14, 44, 26, 21, 6, 94, 33, 40, 2, 94, 89},
		[]int{60, 76, 61, 6, 7, 71, 22, 26, 100, 63, 17, 2, 89, 19, 100, 69}, [][]int{{80, 18}, {99, 27}, {11, 16}, {36, 86}, {98, 80}, {83, 15}, {47, 31}, {7, 70}, {94, 64}, {16, 48}, {33, 41}, {89, 86}, {61, 54}, {25, 40}, {50, 9}, {38, 84}, {30, 77}, {78, 19}, {88, 15}}))
	fmt.Println(maximumSumQueries3([]int{5, 11, 20, 80, 95, 14, 44, 26, 21, 6, 94, 33, 40, 2, 94, 89},
		[]int{60, 76, 61, 6, 7, 71, 22, 26, 100, 63, 17, 2, 89, 19, 100, 69}, [][]int{{80, 18}, {99, 27}, {11, 16}, {36, 86}, {98, 80}, {83, 15}, {47, 31}, {7, 70}, {94, 64}, {16, 48}, {33, 41}, {89, 86}, {61, 54}, {25, 40}, {50, 9}, {38, 84}, {30, 77}, {78, 19}, {88, 15}}))
}
