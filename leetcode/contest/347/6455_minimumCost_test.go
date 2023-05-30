package _347_test

import (
	"fmt"
	"testing"
)

/*
6455. 使所有字符相等的最小成本
https://leetcode.cn/problems/minimum-cost-to-make-all-characters-equal/

s = "010101"
方法：从左到右，枚举每一个位置。
对于位置i来说，如果s[i]!=s[i+1]  那么反转成本小那一段。
这样一次翻转肯定能够让s[i]==s[i+1]
在计算的过程中，不需要真的去做字符翻转的操作。
对于除了s[i-1]和s[i]这两个字符来说的其他字符。
如果翻转前相邻字符是相同的，那么反转后也相同。
如果翻转前相邻字符是不同的，那么反转后也是不同的。

*/

func minimumCost(s string) int64 {
	var n = len(s)
	var ans int64
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans += min(i, n-i)
		}
	}

	return ans

}

func min(args ...int) int64 {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return int64(m)
}

func TestMinimunCost(t *testing.T) {
	fmt.Println(minimumCost("0011"))
}
