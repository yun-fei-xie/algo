package _3_hashtable

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/happy-number/description/

这个题一开始没思路，因为不知道如何应对其中出现的无限循环问题
后查看题解：题目中说了会 无限循环，那么也就是说求和的过程中，sum会重复出现，这对解题很重要！
	于是豁然开朗，可以用map记录出现过的数字
另外解题过程中出现的对数字进行按位拆分（/ %）以后应该会经常遇到
*/

func isHappy(n int) bool {
	record := make(map[int]struct{})
	square := n
	record[square] = struct{}{}
	for square != 1 {
		square = Aux(square)
		if _, found := record[square]; found {
			return false
		}

		record[square] = struct{}{}
	}
	return true
}

func Aux(n int) int {
	sum := 0
	for n != 0 {
		mod := n % 10
		n = n / 10
		sum += mod * mod
	}
	return sum
}

func TestIsHappy(t *testing.T) {
	res1 := isHappy(19)
	res2 := isHappy(2)
	fmt.Println(res1)
	fmt.Println(res2)

}
