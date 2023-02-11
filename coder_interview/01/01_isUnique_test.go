package _1

import (
	"fmt"
	"testing"
)

func isUnique(astr string) bool {
	// 如果遍历结束 有相同的 就返回false
	arr := make([]int, 26)
	for _, char := range astr {
		index := char - 'a'
		arr[index]++
		if arr[index] > 1 {
			return false
		}
	}
	return true
}

func TestIsUnique(t *testing.T) {

	s := "abc"
	res := isUnique(s)
	fmt.Println(res)
}
