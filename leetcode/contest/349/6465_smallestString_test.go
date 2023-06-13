package _49

import (
	"fmt"
	"testing"
)

func smallestString(s string) string {

	//贪心 从左到右开始挑选，直到遇到了字符a(这个已经是最小 如果改变 后面将会变大)
	// 还有一种情况：全是a字符，题目要求必须要换 选择最后一个字符换成z

	sBytes := []byte(s)
	var i = 0
	for ; i < len(sBytes); i++ {
		if sBytes[i] == 'a' {
			continue
		} else {
			break
		}
	}
	var replace bool
	for start := i; start < len(sBytes); start++ {
		if sBytes[start] == 'a' {
			break
		} else {
			sBytes[start] = sBytes[start] - 1
			replace = true
		}
	}

	// 如果之前没有执行替换
	if !replace {
		sBytes[len(sBytes)-1] = 'z'
	}

	return string(sBytes)
}

func TestSmallestString(t *testing.T) {
	fmt.Println(smallestString("leetcode"))
	fmt.Println(smallestString("acbbc"))
	fmt.Println(smallestString("a"))
	fmt.Println(smallestString("aaaa"))

}
