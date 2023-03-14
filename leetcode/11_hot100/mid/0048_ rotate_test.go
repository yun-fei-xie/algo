package mid

/*
https://leetcode.cn/problems/rotate-image/?favorite=2cktkvj
原地旋转一个矩阵,不允许使用另外一个矩阵


解题思路：先转置，再将每一行的元素逆序

1,2,3           1,4,7                        7,4,1
4,5,6   -> 转置  2,5,8   -> 再将每一行元素逆序    8,5,2
7,8,9           3,6,9                        9,6,3


*/

func rotate(matrix [][]int) {

	// 转置
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 逆序

	for i := 0; i < len(matrix); i++ {
		// 第一次用这种数组逆序方式，一行代码
		for j := 0; j < len(matrix[i])/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
}
