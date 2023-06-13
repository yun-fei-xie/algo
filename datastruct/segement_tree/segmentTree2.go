package segement_tree

/*
数组实现线段树，以数组区间和的查询问题为例
空间：所有的节点放在一起，需要的空间上界是4*N。
线段树的数组表示和堆的数组组织方式非常相似。
在线段树的实现中，数据的起始下标从0开始（堆一般从1开始）数据按照层序码放。
*/
type merge func(i, j int) int
type segmentTree struct {
	arr   []int // 用户的原始数据
	tree  []int // 线段树
	merge merge
}

// 初始化
func NewSegmentTree(data []int, m merge) *segmentTree {
	length := len(data)
	arr := make([]int, length)
	copy(arr, data)

	tree := make([]int, 4*length)

	st := &segmentTree{
		arr:   arr,
		tree:  tree,
		merge: m,
	}
	st.buildSegmentTree(0, 0, length-1)
	return st
}

// 构建线段树
func (st *segmentTree) buildSegmentTree(treeIndex int, l int, r int) {
	// 如果区间只有一个元素,递归到底
	if l == r {
		st.tree[treeIndex] = st.arr[l]
		return
	}
	// 否则的话，找出treeIndex的左右子树
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)
	mid := (r-l)/2 + l

	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// 区间查询
func (st *segmentTree) Query(queryL, queryR int) int {
	return st.query(0, 0, len(st.arr)-1, queryL, queryR)
}

// 传入树的根节点和该节点管理的左右区间  传入在这个区间中需要查询的区间
func (st *segmentTree) query(treeIndex, left, right, queryL, queryR int) int {
	if left == queryL && right == queryR {
		return st.tree[treeIndex]
	}
	mid := left + (right-left)/2
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)
	if queryL > mid {
		return st.query(rightTreeIndex, mid+1, right, queryL, queryR)
	}
	if queryR <= mid {
		return st.query(leftTreeIndex, left, mid, queryL, queryR)
	}

	return st.merge(st.query(leftTreeIndex, left, mid, queryL, mid), st.query(rightTreeIndex, mid+1, right, mid+1, queryR))
}

// 单点更新，通过递归找到叶子节点，然后在递归函数返回的时候，将沿途的节点进行融合
// 单点更新可以看做是区间更新，只是这个区间的长度为1，[index ,index]
func (st *segmentTree) Update(index int, val int) {
	st.update(0, 0, len(st.arr)-1, index, val)
}

// 更新操作和build操作本质是一样的
func (st *segmentTree) update(treeIndex int, left int, right int, index int, val int) {
	if left == right && right == index {
		st.tree[treeIndex] = val
		return
	}
	mid := left + (right-left)/2
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	if index <= mid {
		st.update(leftTreeIndex, left, mid, index, val)
		st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
		return
	}
	if index > mid {
		st.update(rightTreeIndex, mid+1, right, index, val)
		st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
	}

}

func (st *segmentTree) leftChild(idx int) int {
	return 2*idx + 1
}
func (st *segmentTree) rightChild(idx int) int {
	return 2*idx + 2
}
