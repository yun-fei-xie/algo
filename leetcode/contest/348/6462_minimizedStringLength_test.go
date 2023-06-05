package _348__test

import (
	"fmt"
	"testing"
)

/*
2716. 最小化字符串长度
https://leetcode.cn/problems/minimize-string-length/

给你一个下标从 0 开始的字符串 s ，重复执行下述操作 任意 次：
在字符串中选出一个下标 i ，并使 c 为字符串下标 i 处的字符。并在 i 左侧（如果有）和 右侧（如果有）各 删除 一个距离 i 最近 的字符 c 。
请你通过执行上述操作任意次，使 s 的长度 最小化 。
返回一个表示 最小化 字符串的长度的整数。

方法：本质是统计字符串中不同字符的个数
*/
func minimizedStringLength(s string) int {
	hashMap := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}
	return len(hashMap)
	//count := 0
	//for _, v := range hashMap {
	//	if v > 1 {
	//		count += v - 1
	//	}
	//}
	//return len(s) - count
}

func TestMinimizedString(t *testing.T) {
	fmt.Println(minimizedStringLength("dddaaa"))
	fmt.Println(minimizedStringLength("aaabc"))
}
