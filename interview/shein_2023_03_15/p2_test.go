package shein_2023_03_15

import (
	"fmt"
	"testing"
)

/*
题目：给你一个字符串，找出所有没有重复字符的长度最长的连续子串
例如：abcdabcd
输出：[abcd,bcda,cdab,dabc,abcd]

思路比较简单的一道题，直接暴力求解。用map记录前面子串中是否出现过重复的字符
*/

func noRepeatSubString(input string) []string {
	maxLength := 0
	records := make(map[int][][]int, 0)
	for i := 0; i < len(input); i++ {
		mp := make(map[uint8]struct{})
		j := i
		for j < len(input) {
			if _, found := mp[input[j]]; !found {
				mp[input[j]] = struct{}{}
				j++
				continue
			} else {
				break
			}
		}
		if j-i > maxLength {
			maxLength = j - i
		}
		if record, found := records[j-i]; found {
			record = append(record, []int{i, j})
			records[j-i] = record
		} else {
			record = [][]int{{i, j}}
			records[j-i] = record
		}
	}

	maxIndexes := records[maxLength]
	ans := make([]string, 0)

	for i := 0; i < len(maxIndexes); i++ {
		ans = append(ans, input[maxIndexes[i][0]:maxIndexes[i][1]])
	}
	return ans
}

func TestNoRepeatSubString(t *testing.T) {
	fmt.Println(noRepeatSubString("abcdabcd"))
}
