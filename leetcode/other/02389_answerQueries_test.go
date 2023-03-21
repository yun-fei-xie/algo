package other

import (
	"fmt"
	"sort"
	"testing"
)

/*
由于元素和有上限，为了能让子序列尽量长，子序列中的元素值越小越好。
把nums从小到大排序后，再从小到大选择尽量多的元素（相当于选择一个前缀），使这些元素的和不超过询问值。

排序+前缀数组
*/
func answerQueries(nums []int, queries []int) []int {
	res := make([]int, len(queries))
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	//找到大于 queries[i]的第一个数的下标，由于下标是从0开始的，这个数的下标正好就是前缀和小于等于 queries[i]的最长前缀的长度。

	for i := 0; i < len(queries); i++ {

		res[i] = sort.SearchInts(nums, queries[i]+1) // 搜索排序数组中值等于queries[i]+1的元素，如果找不到，返回提供的值在数组中的index。
	}
	return res
}

func TestAnswerQueries(t *testing.T) {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
}
