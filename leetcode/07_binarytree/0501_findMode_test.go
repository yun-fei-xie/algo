package _7_binarytree

import (
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/

给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
如果树中有不止一个众数，可以按 任意顺序 返回。

假定 BST 满足如下定义：
结点左子树中所含节点的值 小于等于 当前节点的值
结点右子树中所含节点的值 大于等于 当前节点的值
左子树和右子树都是二叉搜索树

1.把数据放到map[value]-> frequency中，然后在map中进行filter。这种做法消耗掉一定的空间。
2.可以在遍历过程中，记住出现频率最高的val。因为中序遍历是一个有序的过程，相同的元素会排列在一起。
*/

func findMode2(root *TreeNode) []int {
	res := make([]int, 0)
	var base, cnt, maxCnt int
	update := func(val int) {

		if val == base {
			cnt++ // 这里不要直接return了
		} else {
			base, cnt = val, 1
		}
		if cnt == maxCnt {
			res = append(res, val)
		}
		if cnt > maxCnt { // 新的众数出现， 应该将前面的值清空
			res = []int{base}
		}
	}

	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		preOrder(node.Left)
		update(node.Val)
		preOrder(node.Right)
	}

	preOrder(root)
	return res
}

func findMode(root *TreeNode) []int {

	var mp = make(map[int]int)
	var res = make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		_, found := mp[node.Val]
		if found {
			mp[node.Val]++
		} else {
			mp[node.Val] = 1
		}

		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)

	// 从map中找出频率最大的值，这样会比较费空间
	freq := math.MinInt64
	for _, f := range mp {
		if f > freq {
			freq = f
		}
	}
	// 既然不好确定 再遍历一次也无妨。 任何时候，只要有思路，觉得能行，即使不是最优解法 也要先解出来
	for val, f := range mp {
		if f == freq {
			res = append(res, val)
		}
	}
	return res
}

func TestFindMode(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node1 := &TreeNode{
		Val: 2,
	}

	node2 := &TreeNode{
		Val: 2,
	}

	root.Right = node1
	node1.Left = node2

	findMode(root)

}
