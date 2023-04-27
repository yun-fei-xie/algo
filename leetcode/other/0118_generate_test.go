package other

import (
	"fmt"
	"testing"
)

/*
生成杨辉三角
*/
/*
[1],
[1,1],
[1,2,1],
[1,3,3,1],
[1,4,6,4,1]]
*/
func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{1})
	if numRows == 1 {
		return ans
	}
	for i := 1; i < numRows; i++ {
		arr := make([]int, i+1)
		for j := 0; j < (i + 1); j++ {
			var num1, num2 int
			// 每一行只有2个位置需要特殊处理
			if j != 0 {
				num2 = ans[i-1][j-1]
			}
			if j != i {
				num1 = ans[i-1][j]
			}
			arr[j] = num1 + num2
		}
		ans = append(ans, arr)
	}
	return ans
}

func TestGenerate(t *testing.T) {
	fmt.Println(generate(5))
}
