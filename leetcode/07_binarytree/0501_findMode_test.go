package _7_binarytree

import (
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/
*/
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
