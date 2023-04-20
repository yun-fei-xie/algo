package insert_sort

import (
	"fmt"
	"testing"
)

type InsertSort struct {
}

func (ss *InsertSort) sort(arr []int) {

	// 插入排序，假定当前元素前面的元素都是有序的，然后它将会选择一个合适的位置插进入。
	for i := 0; i < len(arr); i++ {
		// 将arr[i]这个元素插入到[0...i-1] 这个区间
		for j := i; j > 0 && (arr)[j] < (arr)[j-1]; j-- {
			(arr)[j-1], (arr)[j] = (arr)[j], (arr)[j-1]
		}
	}

}

func TestInsertSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	ss := InsertSort{}
	ss.sort(arr)
	fmt.Println(arr)
}
