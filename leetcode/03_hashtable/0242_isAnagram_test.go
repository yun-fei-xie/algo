package _3_hashtable

import (
	"fmt"
	"testing"
)

/*
*

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

解法0. 字符串排序，然后比较相当不相等（排序算法nLogn 比较n  时间复杂度n+nLogn -> nLogn 这个推断臆想的）
解法1. 最多有26个字母 数组空间26即可,把字符空间压缩到0-25 然后比较数组中的每一位

但是对于具有更多的字符中类的字符串我应该怎么办？
*/
func isAnagram(s string, t string) bool {
	arrS := make([]int, 26)
	arrT := make([]int, 26)

	for _, charS := range s {
		arrS[charS-'a']++
	}

	for _, charT := range t {
		arrT[charT-'a']++
	}

	for i := 0; i < 26; i++ {
		if arrT[i] != arrS[i] {
			return false
		}
	}

	return true
}

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
	res := isAnagram(s1, s2)
	fmt.Println(res)

}
