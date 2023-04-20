package selectSort

import (
	"fmt"
	"testing"
)

type SelectionSort struct {
}

func (s *SelectionSort) sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		// 寻找从[i...n) 这个区间中最小值索引
		for j := i + 1; j < len(arr); j++ {
			if (arr)[j] < (arr)[i] {
				minIndex = j
			}
			(arr)[i], (arr)[minIndex] = (arr)[minIndex], (arr)[i]
		}
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{10, 9, 8, 7, 6, 4, 3, 2, 1}
	ss := SelectionSort{}
	ss.sort(arr)
	fmt.Println(arr)
}
