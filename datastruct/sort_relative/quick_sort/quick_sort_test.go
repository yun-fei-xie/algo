package quick_sort

import (
	"fmt"
	"testing"
)

type QuickSort struct {
}

func (s *QuickSort) Sort(arr []int) {
	s.sort(arr, 0, len(arr)-1)
}

func (s *QuickSort) sort(arr []int, left, right int) {
	if left > right {
		return
	}
	pivotIndex := s.partition(arr, left, right)

	s.sort(arr, left, pivotIndex-1)
	s.sort(arr, pivotIndex+1, right)

}

// 在[left...right]这个区间中寻找一个pivot，其下标为k , 使得 arr[left , k-1] < arr[k] < arr[k+1 ，right]
func (s *QuickSort) partition(arr []int, left int, right int) int {
	pivot := arr[left]
	i := left + 1
	j := right
	for i <= j {
		if arr[i] < pivot {
			i++
		} else if arr[j] > pivot {
			j--
		} else if arr[i] >= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		} else if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// 循环结束 i会越界一位，使用i-1
	arr[i-1], arr[left] = arr[left], arr[i-1]
	return i - 1
}

func TestQuickSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	ss := QuickSort{}
	ss.Sort(arr)
	fmt.Println(arr)

}
