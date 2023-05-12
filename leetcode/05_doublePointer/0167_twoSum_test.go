package _5_doublePointer

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/

给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。
如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
你所设计的解决方案必须只使用常量级的额外空间。

方法1：对向双指针
2->3->4->6->8
l           r
如果发现nums[l] + nums[r] < target的时候，由于元素的有序性nums[l] + {nums[l+1],...、nums[r-1]}中的任意一个都会
小于target。因此，nums[l]不需要和后面这些不可能能的元素相加再去和target进行比较。它应该向前一位，增大nums[l]。
同理，如果发现nums[l]+nums[r] > target的时候， nums[r] + {nums[l+1]、...、nums[r-1]}中的任意一个都会大于target。
因此，nums[r]不需要和前面这些不可能的元素相加再去和target进行比较。它应该退后，减少nums[r]。

这道题需要注意的地方，题目要求的返回index是元素的位置，而不是索引。
具体可以看题目的输出。位置为1的元素索引为0。
*/
func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	ans := make([]int, 2)
	for left, right := 0, length-1; left < right; {
		sum := numbers[left] + numbers[right]
		if sum == target {
			ans[0] = left + 1
			ans[1] = right + 1
			return ans
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return ans
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
