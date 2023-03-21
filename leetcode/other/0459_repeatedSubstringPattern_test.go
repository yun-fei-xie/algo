package other

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/repeated-substring-pattern/description/
难道倍数不行？
*/

func repeatedSubstringPattern(s string) bool {

	mp := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
	}

	return true

}

func TestRepeatedSubString(t *testing.T) {
	fmt.Println(repeatedSubstringPattern("aba"))

}
