package _5_binarySearch

import (
	"fmt"
	"testing"
)

/*
153. 寻找旋转排序数组中的最小值
https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/

已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。


示例 1：

输入：nums = [3,4,5,1,2]
输出：1
解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
示例 2：

输入：nums = [4,5,6,7,0,1,2]
输出：0
解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。
示例 3：

输入：nums = [11,13,15,17]
输出：11
解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。


题解：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/solutions/698479/xun-zhao-xuan-zhuan-pai-xu-shu-zu-zhong-5irwp/
162问题是峰顶，这个问题是山谷。
对于这个问题来说，[4,5,6,7,0,1,4]  如果最小值的索引是i , 待考察的区间是[left ,right]
那么有nums[i...right-1] < nums[right] && nums[left...i-1] > nums[right] 也就是说nums[i]处在一个最低点。
利用这个性质，先在区间中随便选一个位置(选中点)mid = (left+right)/2
如果nums[mid] < nums[right] ,说明mid这个位置处在上升的这一段。缩小上边界，right=mid (注意，不能直接用mid-1更新right,因为最小值也属于上升的那一段)
如果nums[mid] >=nums[right] ,说明mid这个位置肯定处理下降的那一段，并且不是最小值。（最小值肯定小于right）更新左边界。left=mid+1。
区间中只有一个元素的时候，就是最小值。

*/

func findMin(nums []int) int {

	var length = len(nums)
	var left = 0
	var right = length - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func TestFindMin(t *testing.T) {
	fmt.Println(findMin([]int{1, 2, 3, 4, 5}))
	fmt.Println(findMin([]int{5, 4, 3, 2, 1}))
	fmt.Println(findMin([]int{4, 5, 1, 2, 3}))
	fmt.Println(findMin([]int{10, 20, 30, 40, 1, 2, 3, 4, 5}))

}
