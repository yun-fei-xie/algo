package binary_search

import (
	"fmt"
	"testing"
)

/*

二分查找

1. 从不重复的排序数组中查找值==target的元素
2. 从重复的排序数组中查找


*/

/*
从arr中查找data，如果data不存在，返回其上界的位置
正常的二分查找就是这个位置（如果没有重复数据）
[0,1,2,3,4,6,7,8] 5
返回5 (数字6所在的位置)
*/
func SearchUpperBound(data int, arr []int) int {
	var left, right int
	for left, right = 0, len(arr)-1; left <= right; {
		mid := left + (right-left)/2
		if arr[mid] == data {
			return mid
		} else if arr[mid] > data {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// return left-1 就是下界
	return left // 上界
}

func findRangeCount(arr []int, start int, end int) int {
	leftIdx := binarySearch(arr, start, true)
	rightIdx := binarySearch(arr, end, false)
	return rightIdx - leftIdx
}

func binarySearch(arr []int, target int, isLeft bool) int {
	left, right := 0, len(arr)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			result = mid
			if isLeft {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if isLeft {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return result
}

func TestSearchUpperBound(t *testing.T) {
	//fmt.Println(SearchUpperBound(9, []int{2, 4, 6, 8, 11}))
	//fmt.Println(SearchUpperBound(10, []int{2, 4, 6, 8, 11}) - 1)

	fmt.Println(findRangeCount([]int{2, 10}, 1, 9))
}
