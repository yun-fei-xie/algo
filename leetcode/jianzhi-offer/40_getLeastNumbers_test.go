package jianzhi_offer

import (
	"fmt"
	"testing"
)

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
*/

/*
方法1,排序后返回前k个数字 使用内置api->不是考点
方法2，堆排序 ->相对于快速排序 复杂度要高一些
方法3，快速排序 -> 实现这种方法
*/
func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, arr[i])
	}
	return ans
}

func quickSort(arr []int, left, right int) {
	if left > right {
		return
	}
	pivotIndex := partition(arr, left, right)
	quickSort(arr, left, pivotIndex-1)
	quickSort(arr, pivotIndex+1, right)
}

/*
partition 对arr[left,right]进行partition操作，返回pivot
*/
func partition(arr []int, left int, right int) int {
	if left >= right {
		return left
	}
	l := left + 1
	r := right
	num := arr[left]
	for l <= r {
		if arr[r] > num {
			r--
		} else if arr[l] < num {
			l++
		} else if arr[r] <= num {
			arr[r], arr[l] = arr[l], arr[r]
			r--
		} else if arr[l] >= num {
			arr[r], arr[l] = arr[l], arr[r]
			l++
		}
	}
	arr[left], arr[l-1] = arr[l-1], arr[left]
	return l - 1
}
func TestGetLeastNumbers(t *testing.T) {
	fmt.Println(getLeastNumbers([]int{4, 5, 1, 6, 2, 7, 3, 8}, 3))
}
