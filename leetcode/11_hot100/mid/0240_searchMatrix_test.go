package mid

/*

https://leetcode.cn/problems/search-a-2d-matrix-ii/description/?favorite=2cktkvj
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

解法1：一行一行进行二分查找，结合有序的性质，某些行可以提前结束。比如，当前行第一元素大于target,直接返回，后面行也不需要考虑。

解法2：会发现每次都是向左数字会变小，向下数字会变大，有点和二分查找树相似。二分查找树的话，是向左数字变小，向右数字变大。
所以我们可以把 target 和当前值比较。
如果 target 的值大于当前值，那么就向下走。
如果 target 的值小于当前值，那么就向左走。
如果相等的话，直接返回 true 。



*/

func searchMatrixS1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var res bool

	for i := 0; i < len(matrix); i++ {
		arr := matrix[i]
		if arr[0] > target {
			return res
		}

		left := 0
		right := len(arr) - 1
		for left <= right {
			mid := left + (right-left)>>1

			if arr[mid] == target {
				return true
			} else if arr[mid] > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return res
}

/*
方法2：从右上角开始进行搜索，非常像二叉搜索树
*/
func searchMatrixS2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := 0
	col := len(matrix[0]) - 1
	// 从右上角开始
	for row < len(matrix) && col >= 0 { // 临界条件不要超出矩阵的边界
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
