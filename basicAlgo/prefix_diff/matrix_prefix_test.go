package prefix_diff

import (
	"fmt"
	"testing"
)

/*
矩阵前缀和

原始矩阵
1 2 4 3
5 1 2 4
6 3 5 9

求解前缀矩阵(前缀矩阵就是从（0，0） 这个点到当前点的所有元素的和)
例如，prefix[i][j] =  sum(arr[x][y]) (0<=x<=i , 0<=y<=j)

1  3  7  10
6  9  15 22
12 18 29 45
如何求解呢？
prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + matrix[i][j]

应用：求解子矩阵的和  求解子矩阵的和有些差分的意思
15 , 22
29 , 45 这个子矩阵的和

也就是（1，2） -> （2，3）
如何求解这个范围内的子矩阵的和呢？
用45为右下角前缀和-以10为右下角的前缀和-以18为右下角的前缀和 + 以3为右下角的前缀和
也就是：
prefix[2][3] - prefix[0][3] - prefix[2][1] + prefix[0][1]

更加一般性：从（x1 , y1）-> (x2 , y2)
prefix[x2][y2] - prefix[x1-1][y2] - prefix[x2][y1-1] + prefix[x1-1][y1-1]
*/
func subMatrix(prefix [][]int, x1, y1, x2, y2 int) int {

	var prefix1, prefix2, prefix3 int
	if x1-1 >= 0 {
		prefix1 = prefix[x1-1][y2]
	}

	if y1-1 >= 0 {
		prefix2 = prefix[x2][y1-1]
	}

	if x1-1 >= 0 && y1-1 >= 0 {
		prefix3 = prefix[x1-1][y1-1]
	}

	return prefix[x2][y2] - prefix1 - prefix2 + prefix3
}

func matrixPrefix(matrix [][]int) [][]int {

	row := len(matrix)
	col := len(matrix[0])
	var prefix [][]int
	for i := 0; i < row; i++ {
		prefix = append(prefix, make([]int, col))
	}
	// init
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			var prefix1, prefix2, prefix3 int

			if i-1 >= 0 && j >= 0 {
				prefix1 = prefix[i-1][j]
			}

			if i >= 0 && j-1 >= 0 {
				prefix2 = prefix[i][j-1]
			}

			if i-1 >= 0 && j-1 >= 0 {
				prefix3 = prefix[i-1][j-1]
			}

			prefix[i][j] = prefix1 + prefix2 - prefix3 + matrix[i][j]

		}
	}

	return prefix
}

func TestMatrixPrefix(t *testing.T) {

	var matrix = [][]int{{1, 2, 4, 3}, {5, 1, 2, 4}, {6, 3, 5, 9}}
	prefix := matrixPrefix(matrix)
	for i := 0; i < len(prefix); i++ {
		fmt.Println(prefix[i])
	}

	subMatrixSum := subMatrix(prefix, 1, 2, 2, 3)
	fmt.Println(subMatrixSum)

}
