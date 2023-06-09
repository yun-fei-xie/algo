package st

import (
	"fmt"
	"testing"
)

/*
st表的简单示例，给定一个数组，查询数组中的区间最大值
*/
type sTable struct {
	st [][]int
}

func newST(arr []int) *sTable {
	maxLength := len(arr)
	maxIndex := len(arr) - 1
	s := make([][]int, maxIndex+1) // [0...maxIndex]
	for i := 0; i < len(s); i++ {
		s[i] = make([]int, maxLength+1) //[1...maxLength]
	}

	// 填写表格
	for i := 0; i < len(arr); i++ {
		s[i][1] = arr[i]
	}
	for i := 0; i < len(arr); i++ {
		for j := 2; j <= maxLength && (i+j-1) < len(arr); j++ {
			s[i][j] = max(s[i][j-1], arr[i+j-1])
		}
	}
	return &sTable{st: s}
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
查询操作
*/
func (st *sTable) query(queries [][]int) (ans []int) {
	for _, query := range queries {
		left := query[0]
		length := query[1] - query[0] + 1
		ans = append(ans, st.st[left][length])
	}
	return ans
}

func TestSt(t *testing.T) {
	arr := []int{5, 3, 7, 2, 12, 1} //[0...5]
	st := newST(arr)
	queries := [][]int{{0, 1}, {0, 4}, {1, 3}, {1, 2}, {3, 5}, {2, 5}}
	ans := st.query(queries)
	fmt.Println(ans)
}
