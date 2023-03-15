package _7_binarytree

import "container/list"

/*
https://leetcode.cn/problems/binary-tree-right-side-view/
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

咋一看以为只需要沿着右子树遍历即可。实际上是不对的（想象一下全是左子树的情况）

解法1（递归）: 换一个思考方式，它其实需要的是每一层最右边的那个元素（层序遍历）收集数据
然后把每一层的最后一位收集起来。主要注意的是，在递归层序遍历的时候，同一层的元素不是连续遍历的。而是使用depth对当前遍历到的元素进行层数标记。

解法2（迭代）: 用队列遍历的时候判断是否遍历到这一层的最右边的元素，如果是则放入数组

*/
/*
递归版本
*/
func rightSideView1(root *TreeNode) []int {
	values := make([]int, 0)

	data := make([][]int, 0)

	var levelOrder func(node *TreeNode, depth int)
	levelOrder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(data) == depth {
			data = append(data, []int{})
		}

		data[depth] = append(data[depth], node.Val)
		levelOrder(node.Left, depth+1)
		levelOrder(node.Right, depth+1)
	}
	levelOrder(root, 0)
	for i := 0; i < len(data); i++ {
		levelData := data[i]
		values = append(values, levelData[len(levelData)-1])
	}
	return values
}

/*
非递归版本
*/
func rightSideView(root *TreeNode) []int {
	values := make([]int, 0)
	if root == nil {
		return values
	}
	queue := list.New()
	levelCount := 1
	queue.PushBack(root)
	for queue.Len() > 0 {

		nextLevelCount := 0
		for i := 0; i < levelCount; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if i == levelCount-1 {
				values = append(values, node.Val)
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
				nextLevelCount++
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
				nextLevelCount++
			}
		}
		levelCount = nextLevelCount
	}
	return values
}

/*
想当然版本（错误版本）
*/
func rightSideViewErr(root *TreeNode) []int {

	values := make([]int, 0)
	if root == nil {
		return values
	}

	cur := root

	for cur != nil {

		values = append(values, cur.Val)
		cur = cur.Right

	}
	return values

}
