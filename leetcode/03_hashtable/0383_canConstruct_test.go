package _3_hashtable

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/ransom-note/description/

map还是有些浪费了，应该使用26位长度的数组进行记录
*/
func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[int32]int)
	m2 := make(map[int32]int)
	for _, char1 := range ransomNote {
		m1[char1-'a']++
	}
	for _, char2 := range magazine {
		m2[char2-'a']++
	}

	for key, value := range m1 {
		if v, found := m2[key]; !found || v < value {
			return false
		}
	}
	return true
}
func canConstruct2(ransomNote string, magazine string) bool {
	m1 := [26]int{}
	m2 := [26]int{}

	for _, char1 := range ransomNote {
		m1[char1-'a']++
	}
	for _, char2 := range magazine {
		m2[char2-'a']++
	}
	for i := 0; i < len(m1); i++ {
		if m2[i] < m1[i] {
			return false
		}
	}
	return true
}

func TestCanConstruct(t *testing.T) {
	ransomNote := "aa"
	magazine := "aab"

	res := canConstruct2(ransomNote, magazine)
	fmt.Println(res)
}
