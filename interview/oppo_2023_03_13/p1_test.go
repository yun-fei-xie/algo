package oppo_2023_03_13

import (
	"fmt"
	"sort"
	"testing"
)

/*
[3,4]
43
*/
func maxDigit(digits []int) int {
	// write code here

	sort.Slice(digits, func(i, j int) bool {
		if digits[i] > digits[j] {
			return true
		}
		return false
	})

	res := 0
	for i := 0; i < len(digits); i++ {
		res = res*10 + digits[i]
	}

	return res
}

func TestMaxDigit2(t *testing.T) {
	res := maxDigit([]int{1, 2, 3, 4})
	fmt.Println(res)
}

func TestMaxDigit(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] > arr[j] {
			return true
		}
		return false
	})

	fmt.Println(arr)
}
