package _7_binarytree

import "container/list"

/*
 */
func levelOrder(root *TreeNode) [][]int {

	values := make([][]int, 0)
	queue := list.New()
	if root == nil {
		return values
	}
	queue.PushBack(root)
	levelCount := 1
	for queue.Len() != 0 {

		// 1.取出当前层  2.放入数组 3.放入下一层  ？（怎么知道哪些元素是当前层 用一个变量标记 levelCount）
		nextLevelCount := 0
		var levelValues = make([]int, 0)
		for i := 0; i < levelCount; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			levelValues = append(levelValues, node.Val)
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
		values = append(values, levelValues)
	}
	return values
}

/*
*
层序遍历 递归写法
每走到一层，就把当前层的数据放入对应的数组中
层数用depth参数进行标记
*/
func levelOrderRec(root *TreeNode) [][]int {

	values := [][]int{}
	depth := 0
	var levelorder func(node *TreeNode, depth int)
	levelorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(values) == depth {
			values = append(values, []int{})
		}
		values[depth] = append(values[depth], node.Val)
		levelorder(node.Left, depth+1)
		levelorder(node.Right, depth+1)
	}
	levelorder(root, depth)

	return values
}
