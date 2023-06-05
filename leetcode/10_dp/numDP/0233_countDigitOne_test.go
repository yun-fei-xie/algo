package numDP__test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
233. 数字 1 的个数
https://leetcode.cn/problems/number-of-digit-one/
给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
13->6 {1 ,10 , 11 , 12, 13} 一共有6个1

方法：枚举 数位dp
1.将整数转为字符串，同样是构造数字。
2.将[0...i]中已经拿到的1的个数放入递归函数中，向下传递。
3.递归结束的时候，将countOne变量返回给上一层。
*/
func countDigitOne(n int) int {
	s := strconv.Itoa(n)

	var numDp func(i int, countOne int, hasLimit bool) (res int)
	numDp = func(i int, countOne int, hasLimit bool) (res int) {
		if i == len(s) {
			return countOne
		}
		if hasLimit {
			for j := 0; j <= int(s[i]-'0'); j++ {
				if j == 1 {
					res += numDp(i+1, countOne+1, hasLimit && j == int(s[i]-'0'))
				} else {
					res += numDp(i+1, countOne, hasLimit && j == int(s[i]-'0'))
				}
			}
		} else {
			for j := 0; j <= 9; j++ {
				if j == 1 {
					res += numDp(i+1, countOne+1, hasLimit && j == int(s[i]-'0'))
				} else {
					res += numDp(i+1, countOne, hasLimit && j == int(s[i]-'0'))
				}
			}
		}
		return res
	}
	return numDp(0, 0, true)
}

/*
记忆化
*/

func countDigitOne2(n int) int {
	s := strconv.Itoa(n)
	// mem数组需要记忆从i这个位置开始，[i...len(s))，能够出现多少个1.
	// 同样沿用之前

	mem := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		mem[i] = make([]int, len(s)+1)
		for j := 0; j < len(s)+1; j++ {
			mem[i][j] = -1
		}

	}

	// 递归函数是在构造数字，如果构造成功，它会返回这个数字中1的个数。
	var numDp func(i int, countOne int, hasLimit bool) (res int)
	numDp = func(i int, countOne int, hasLimit bool) (res int) {
		if i == len(s) {
			return countOne
		}

		if !hasLimit {
			if mem[i][countOne] != -1 {
				return mem[i][countOne]
			}
			defer func() {
				mem[i][countOne] = res
			}()
		}

		if hasLimit {
			for j := 0; j <= int(s[i]-'0'); j++ {
				if j == 1 {
					res += numDp(i+1, countOne+1, hasLimit && j == int(s[i]-'0'))
				} else {
					res += numDp(i+1, countOne, hasLimit && j == int(s[i]-'0'))
				}
			}
		} else {
			for j := 0; j <= 9; j++ {
				if j == 1 {
					res += numDp(i+1, countOne+1, hasLimit && j == int(s[i]-'0'))
				} else {
					res += numDp(i+1, countOne, hasLimit && j == int(s[i]-'0'))
				}
			}
		}
		return res
	}
	return numDp(0, 0, true)
}

func TestCountDigitOne(t *testing.T) {
	fmt.Println(countDigitOne(13))
	fmt.Println(countDigitOne(0))
}
