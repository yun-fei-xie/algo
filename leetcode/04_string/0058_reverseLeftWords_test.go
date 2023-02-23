package _4_string

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

感觉需要在k上做文章
如果长度为6 左旋转k 相当于右旋转 length-k  而k需要做处理 k = k %length

于是就有 k = k %length   m = length - k  向右旋转m
如果之前的位置位i 新的位置位 (i+m）%（length）
注意边界，如果length = 7 ,i=6 向右1位-> 应该到0
（6+1）% 7 = 0

数组下标旋转 首先想到的就是取模


*/

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	res := make([]byte, len(b))
	n = n % len(b)
	m := len(b) - n

	for i := 0; i < len(b); i++ {

		res[(i+m)%len(b)] = b[i]
	}
	return string(res)
}

func TestReverseLeftWords(t *testing.T) {
	s := "abcdefg"
	k := 2

	res := reverseLeftWords(s, k)
	fmt.Println(res)
}
