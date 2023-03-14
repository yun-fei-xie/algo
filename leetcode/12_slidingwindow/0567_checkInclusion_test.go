package _2_slidingwindow

import (
	"fmt"
	"reflect"
	"testing"
)

/*
https://leetcode.cn/problems/permutation-in-string/
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

滑动窗口找子串，涉及排列用map进行比较
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
*/

func checkInclusion(s1 string, s2 string) bool {
	mp1 := make(map[uint8]int)
	for i := 0; i < len(s1); i++ {
		mp1[s1[i]]++
	}
	res := false
	mp2 := make(map[uint8]int)
	left := 0
	right := 0
	for right < len(s2) {
		if right < len(s1)-1 { //init
			mp2[s2[right]]++
			right++
		} else {
			mp2[s2[right]]++
			if reflect.DeepEqual(mp1, mp2) {
				res = true
			}
			mp2[s2[left]]--
			if mp2[s2[left]] == 0 {
				delete(mp2, s2[left])
			}

			left++
			right++
		}
	}
	return res
}

func TestCheckInclusion(t *testing.T) {
	res := checkInclusion("ab", "eidbaooo")
	fmt.Println(res)

}
