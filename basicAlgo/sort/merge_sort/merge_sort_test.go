package merge_sort

import (
	"fmt"
	"testing"
)

type MergeSort struct {
}

/*
归并排序使用了递归
先sort再merge,以中点为界限
*/
func (ss *MergeSort) Sort(arr []int) {
	ss.sort(arr, 0, len(arr)-1)
}

func (ss *MergeSort) sort(arr []int, left int, right int) {
	if left >= right { // 区间中只有一个元素或者没有元素->已经有序
		return
	}
	mid := (left + right) / 2
	ss.sort(arr, left, mid)
	ss.sort(arr, mid+1, right)
	ss.merge(arr, left, mid, right)
}

func (ss *MergeSort) merge(arr []int, left int, mid int, right int) {
	auxArr := make([]int, right-left+1)
	copy(auxArr, arr[left:right+1]) // arr[left...right+1)
	var i = left
	var j = mid + 1

	for k := left; k <= right; k++ {
		if i > mid { // 越界的判断必须放在前面
			arr[k] = auxArr[j-left]
			j++
		} else if j > right { // 越界的判断必须放在前面
			arr[k] = auxArr[i-left]
			i++
		} else if auxArr[i-left] < auxArr[j-left] {
			arr[k] = auxArr[i-left]
			i++
		} else if auxArr[i-left] >= auxArr[j-left] {
			arr[k] = auxArr[j-left]
			j++
		}
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	ss := MergeSort{}
	ss.Sort(arr)
	fmt.Println(arr)

}
