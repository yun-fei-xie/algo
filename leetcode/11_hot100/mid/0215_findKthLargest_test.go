package mid

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
和347有点像，基于快速排序的思想

快排核心思想就是分治和分区，我们可以利用分区的思想，来解答开篇的问题：O(n) 时间复杂度内求无序数组中的第 K 大元素。
比如，4， 2， 5， 12， 3 这样一组数据，第 3 大元素就是 4。
我们选择数组区间 A[0...n-1]的最后一个元素 A[n-1]作为 pivot，对数组 A[0...n-1]原地分区，这样数组就分成了三部分，A[0...p-1]、A[p]、A[p+1...n-1]。
如果 p+1=K，那 A[p]就是要求解的元素；如果 K>p+1, 说明第 K 大元素出现在 A[p+1...n-1]区间，我们再按照上面的思路递归地在 A[p+1...n-1]这个区间内查找。
同理，如果 K<p+1，那我们就在 A[0...p-1]区间查找。

上面这种找法，针对没有重复元素可行。但是如果有重复元素就不行。

*/

// 这个解错了是因为数组有重复元素
func findKthLargest(nums []int, k int) int {
	var res = -1
	var quickSearchK func(arr []int, left int, right int)
	quickSearchK = func(arr []int, left int, right int) {
		pivot := partition(arr, left, right) // 拿到的是相对于整个数组的索引
		if pivot+1 == k {
			res = arr[pivot]
		} else if pivot+1 > k {
			quickSearchK(arr, left, pivot-1)
		} else { // pivot +1 < k
			quickSearchK(arr, pivot+1, right)
		}
	}

	quickSearchK(nums, 0, len(nums)-1)

	return res
}

/*
[left , right] 闭区间
*/
func partition(arr []int, left int, right int) int {

	pivot := arr[left]
	begin := left + 1
	end := right
	for begin <= end {
		if arr[begin] <= pivot {
			begin++
		} else if arr[end] > pivot {
			end--
		} else if arr[begin] > pivot {
			arr[begin], arr[end] = arr[end], arr[begin]
			end--
		} else if arr[end] < pivot {
			arr[begin], arr[end] = arr[end], arr[begin]
			begin++
		}
	}
	arr[left], arr[begin-1] = arr[begin-1], arr[left]
	return begin - 1
}

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
