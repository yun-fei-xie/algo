package mid

/*
https://leetcode.cn/problems/path-sum-iii/?favorite=2cktkvj

给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

注意：这里的路径只能是向下。难度降低，只需要走左右子树的一个分支即可。不需要考虑^ 这种类型的情况.

暴力解法：对于每个节点，都尝试以该节点为根节点，向下进行一次搜索。


前缀和：两节点间的路径和 = 两节点的前缀和之差

重复路径-> 下图中，
    1
   /
  0
 /
2

当来到一个节点currSum(前缀和) 我想知道祖先节点有没有前缀和等于 targetSum-currSum的。如果有的话，便是找到一个解。
怎么找呢？从上到下进行遍历的时候，将前缀和保存在map中。并且需要再回退时，从map中删除当前节点的前缀和。
这是因为，我们只能在同一条路径上进行寻找。



这个题和网宿笔试的25分大题有点像，那个题是求二叉树路径最大值。
那个题应该如何进行遍历？
*/

/*
前缀和解法
*/
func pathSum2(root *TreeNode, targetSum int) int {
	res := 0

	//	prefixTable := make(map[int]int, 0) // 前缀和key->出现的次数
	prefixTable := map[int]int{0: 1} // 为什么需要这个？

	var preOrderTravels func(node *TreeNode, prefixSum int)
	preOrderTravels = func(node *TreeNode, prefixSum int) {
		if node == nil {
			return
		}

		newPrefixSum := prefixSum + node.Val

		if cnt, found := prefixTable[newPrefixSum-targetSum]; found {
			res += cnt
		}

		// 更新前缀和
		if cnt, found := prefixTable[newPrefixSum]; found {
			cnt++
			prefixTable[newPrefixSum] = cnt
		} else {
			prefixTable[newPrefixSum] = 1
		}
		// 向下递归
		preOrderTravels(node.Left, newPrefixSum)
		preOrderTravels(node.Right, newPrefixSum)

		// 回退的时候，删除当前的前缀和
		prefixTable[newPrefixSum]--
		if prefixTable[newPrefixSum] == 0 {
			delete(prefixTable, newPrefixSum)
		}
	}

	preOrderTravels(root, 0)

	return res
}

/*
这样解答为什么不对呢？
*/
func pathSum(root *TreeNode, targetSum int) int {
	res := 0

	var preOrderTravels func(node *TreeNode)
	var search func(node *TreeNode, sum int)
	preOrderTravels = func(node *TreeNode) {
		if node == nil {
			return
		}
		search(node, node.Val)
		preOrderTravels(node.Left)
		preOrderTravels(node.Right)
	}

	search = func(node *TreeNode, sum int) {
		if sum == targetSum || node == nil {
			if sum == targetSum {
				res += 1
			}
			return
		}
		if node.Left != nil {

			search(node.Left, sum+node.Left.Val)
		}
		if node.Right != nil {

			search(node.Right, sum+node.Right.Val)
		}
	}

	preOrderTravels(root)

	return res
}
