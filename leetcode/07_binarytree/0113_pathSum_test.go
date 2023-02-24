package _7_binarytree

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
有很多点需要注意：
1.slice的指针传递修改数组
2.回溯过程中的路径删除
3.slice的copy
*/

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	pathSumAux(root, targetSum, new([]int), &res)
	return res
}

func pathSumAux(node *TreeNode, target int, path *[]int, result *[][]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {

		curTarget := target - node.Val
		*path = append(*path, node.Val)

		if curTarget == 0 {
			curPath := make([]int, len(*path))
			copy(curPath, *path)
			*result = append(*result, curPath)
		}
		return
	}

	*path = append(*path, node.Val)
	curTarget := target - node.Val

	if node.Left != nil {
		pathSumAux(node.Left, curTarget, path, result)
		*path = (*path)[:len(*path)-1]
	}

	if node.Right != nil {

		pathSumAux(node.Right, curTarget, path, result)
		*path = (*path)[:len(*path)-1]
	}

}

// 在参数中修改数组
func TestSlice(t *testing.T) {

	a := []int{0, 1, 2, 3, 4}
	a = a[:len(a)]
	fmt.Println(a)

	//f := func(arr *[]int) {
	//	*arr = append(*arr, 100)
	//}
	//fmt.Println(a)
	//f(&a)
	//fmt.Println(a)

}
