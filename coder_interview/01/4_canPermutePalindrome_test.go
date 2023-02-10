package _1

import (
	"fmt"
	"testing"
)

/**
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
回文串不一定是字典当中的单词。
示例1：
输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）

*/

/*
*
题解：如果某个字母出现的次数是奇数个 那么不可能是回文串
（除非这个字母放在字符串的中间）
所以，如果有2个字母，都是奇数个数，肯定不能构成回文

ps : 字符串中可能有非字母 因此,不能使用

	arr:=make([]int , 26) -> arr['ch' - 'a']
*/
func canPermutePalindrome(s string) bool {
	m := make(map[int32]int, 0)
	for _, ch := range s {
		m[ch]++
	}
	var count = 0
	for _, num := range m {
		if num%2 != 0 {
			count++
			if count >= 2 {
				return false
			}
		}
	}
	return true
}

func TestCanPermutePalindrome(t *testing.T) {

	res := canPermutePalindrome("taoa")
	fmt.Println(res)

}
