package bubble_sort

import (
	"fmt"
	"testing"
)

type BubbleSort struct {
}

/*
冒泡排序思想：
从左到右，相邻的两个元素进行比较，把大的元素交换到索引大的位置。
*/
func (ss *BubbleSort) sort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	ss := BubbleSort{}
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	ss.sort(arr)
	fmt.Println(arr)

}
