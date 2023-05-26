package segement_tree

/*
链表实现线段树
统计有序数组中，给定区间大小[low ,high]中，数组中包含多少个元素
*/

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
