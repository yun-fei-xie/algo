package _4_kd_tree

import (
	"fmt"
	"testing"
)

/*
使用bst做一维空间range search
测试方式，给你一个数组arr=[...若干点]
给定若干范围查询query={{lo,high},{...}}
返回区间中点的个数，和对应的点是哪些

我们这棵树可以支持存储重复的元素。在node中用一个额外的变量记录val出现的次数。

*/

type bstNode struct {
	val   int
	cnt   int
	left  *bstNode
	right *bstNode
}

type BST struct {
	root *bstNode
	size int
}

func NewBst(arr []int) *BST {
	bst := &BST{
		root: nil,
		size: len(arr),
	}
	for i := 0; i < len(arr); i++ {
		bst.root = bst.addValue(bst.root, arr[i])
	}
	return bst
}

func (bst *BST) addValue(root *bstNode, val int) *bstNode {
	if root == nil {
		return &bstNode{
			val: val,
			cnt: 1,
		}
	}
	if val == root.val {
		root.cnt++
		return root
	} else if val > root.val {
		root.right = bst.addValue(root.right, val)
		return root
	} else {
		root.left = bst.addValue(root.left, val)
		return root
	}
}

// 查找闭区间[lo,hi]中有多少元素，有哪些元素
// 查找策略：考虑当前节点的左子树、考虑当前节点的右子树、考虑当前节点
// 假设当前节点的值为val:
// 1.如果当前节点的左子树小于lo或者为空，那么将不会去左子树查询，否则递归去左子树查询
// 2.如果当前节点的右子树大于hi或者为空，那么将不会去右子树查询，否则递归去右子树查询
// 3.查看当期那节点是否在闭区间中。
func (bst *BST) QueryRange(lo int, hi int) (cnt int, nodes []int) {
	if bst.root == nil {
		return 0, nil
	}
	var queryRange func(root *bstNode, l int, r int)
	queryRange = func(root *bstNode, l int, r int) {
		if root.val >= l && root.val <= r {
			cnt += root.cnt
			for i := 0; i < root.cnt; i++ {
				nodes = append(nodes, root.val)
			}
		}
		// 如果当前节点的值小于等于l，那么左子树必然都小于l 不用看左子树。否则去左子树中递归查询。
		if root.val > l && root.left != nil {
			queryRange(root.left, l, r)
		}
		// 如果当前节点的值大于等于r，那么当前节点的右子树都大于r，不用看右子树。否则去右子树中递归查询。
		if root.val < r && root.right != nil {
			queryRange(root.right, l, r)
		}

	}
	queryRange(bst.root, lo, hi)
	return cnt, nodes
}

func TestBSTRangeQuery(t *testing.T) {
	arr := []int{3, 5, 17, 7, 15, 1, 9, 11, 13}
	bst := NewBst(arr)
	cnt, nodes := bst.QueryRange(2, 10)
	fmt.Println("cnt nodes ", cnt, nodes)
}
