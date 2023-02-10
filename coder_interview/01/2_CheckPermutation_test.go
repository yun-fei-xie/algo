package _1

/**
给定两个由小写字母组成的字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false

*/
import (
	"fmt"
	"testing"
)

func CheckPermutation(s1 string, s2 string) bool {
	if s1 == "" && s2 == "" {
		return true
	}

	arr1 := make([]int, 26)
	arr2 := make([]int, 26)

	for _, char1 := range s1 {
		arr1[char1-'a']++
	}

	for _, char2 := range s2 {
		arr2[char2-'a']++
	}

	for index, val := range arr1 {
		if arr2[index] != val {
			return false
		}
	}

	return true
}

func Test(t *testing.T) {
	s1 := "abc"
	s2 := "bad"
	res := CheckPermutation(s1, s2)
	fmt.Println(res)
}
