package main

import (
	"fmt"
	"sort"
)

/*
Description
C 公司对一个即将上线的业务进行了测试，为了更好的评估该业务测试期间的运行状态，C 公司的 Alice 同学打算写一个工具来对该业务的 Access Log 进行处理，她希望能通过这个工具知道任意一个时间段内 (闭区间，精确到秒) 该业务有多少次访问。
为了方便处理，机智的 Alice 同学写了个脚本将 Access Log 简化为了每条日志只有一个时间戳数据，即简化后的 Access Log 是一个 N 行的文本，每行是一个整数时间戳，代表该时刻有一次访问，如下图所示:

1564900123
1564934135
1564934132
1564934666
1564931024
但是 Alice 同学太忙了，所以需要请你帮她实现这个问题。

*/

func main1() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var N, M int
		fmt.Scan(&N)

		logs := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&logs[j])
		}

		sort.Ints(logs)
		segTree := BuildSegmentTree(logs)

		fmt.Scan(&M)
		for j := 0; j < M; j++ {
			var b, e int
			fmt.Scan(&b, &e)
			count := segTree.Query(b, e)
			fmt.Printf("%d", count)
			if j != M-1 {
				fmt.Println()
			}
		}
	}
}

func mainP1() {

	var t int
	fmt.Scan(&t)

	var query func(arr []uint64, left uint64, right uint64) (count uint64)
	query = func(arr []uint64, left uint64, right uint64) (count uint64) {
		for i := 0; i < len(arr); i++ {
			if arr[i] >= left && arr[i] <= right {
				count++
			}
		}
		return count
	}

	for i := 0; i < t; i++ {
		var N, M int
		fmt.Scan(&N)

		logs := make([]uint64, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&logs[j])
		}
		fmt.Scan(&M)
		for j := 0; j < M; j++ {
			var b, e uint64
			fmt.Scan(&b, &e)
			count := query(logs, b, e)
			fmt.Printf("%d", count)
			if j != M-1 {
				fmt.Println()
			}
		}
	}

}

// 线段树求解

type SegmentTree struct {
	root *segmentNode
}

// SegmentNode 线段树节点：leftVal表示区间的左端点值 rightVal表示区间的右端点值 count表示arr[left,right]中有多少元素
type segmentNode struct {
	leftVal    int
	rightVal   int
	count      int
	leftChild  *segmentNode
	rightChild *segmentNode
}

func newSegmentNode(leftVal int, rightVal int, count int, leftChild *segmentNode, rightChild *segmentNode) *segmentNode {
	return &segmentNode{
		leftVal:    leftVal,
		rightVal:   rightVal,
		count:      count,
		leftChild:  leftChild,
		rightChild: rightChild,
	}
}

// BuildSegmentTree 构建线段树是一个递归的过程
func BuildSegmentTree(arr []int) *SegmentTree {
	root := buildTree(arr[0], arr[len(arr)-1], arr)
	return &SegmentTree{root: root}
}

// 辅助函数，构建数组arr下标索引从[left...right]这段区间的线段树
func buildTree(left int, right int, arr []int) (root *segmentNode) {
	if left == right {
		var count int
		if binarySearch(arr, left) {
			count = 1
		}
		return newSegmentNode(left, right, count, nil, nil)
	}
	mid := left + (right-left)/2
	leftChild := buildTree(left, mid, arr)
	rightChild := buildTree(mid+1, right, arr)

	root = newSegmentNode(left, right, leftChild.count+rightChild.count, leftChild, rightChild)
	return root
}

func binarySearch(arr []int, target int) bool {
	for left, right := 0, len(arr)-1; left <= right; {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// Query 区间查询
func (seg *SegmentTree) Query(left int, right int) (count int) {

	if left < seg.root.leftVal {
		left = seg.root.leftVal
	}
	if right > seg.root.rightVal {
		right = seg.root.rightVal
	}

	if left > seg.root.rightVal || right < seg.root.leftVal {
		return 0
	}
	return query(seg.root, left, right)
}
func query(root *segmentNode, left int, right int) (count int) {

	//fmt.Printf("调用query 当前节点左右区间[%d  %d] ,当前查询区间[%d  %d]\n", root.leftVal, root.rightVal, left, right)

	if left == root.leftVal && right == root.rightVal {
		return root.count
	}
	mid := root.leftVal + (root.rightVal-root.leftVal)/2

	if left > mid {
		return query(root.rightChild, left, right)
	} else if right <= mid {
		return query(root.leftChild, left, right)
	} else {
		return query(root.leftChild, left, mid) + query(root.rightChild, mid+1, right)
	}
}
