package mid

import (
	"fmt"
	"math/rand"
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

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	targetIndex := length - k
	left := 0
	right := length - 1
	for {
		pivotIndex := partition(nums, left, right)
		if pivotIndex == targetIndex {
			return nums[pivotIndex]
		} else if pivotIndex > targetIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}

}

/*
[left , right] 闭区间
*/
func partition(arr []int, left int, right int) int {
	//pivotNum := arr[left]
	//随机 让索引落到[left,right]
	randIndex := rand.Intn(right-left+1) + left
	arr[randIndex], arr[left] = arr[left], arr[randIndex]

	pivotNum := arr[left]
	i := left + 1
	j := right
	for i <= j {
		if arr[j] > pivotNum {
			j--
		} else if arr[i] <= pivotNum {
			i++
		} else if arr[j] <= pivotNum {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		} else if arr[i] > pivotNum {
			arr[j], arr[i] = arr[i], arr[j]
			j--
		}
	}
	arr[i-1], arr[left] = arr[left], arr[i-1]
	return i - 1
}

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{1}, 1))
}
