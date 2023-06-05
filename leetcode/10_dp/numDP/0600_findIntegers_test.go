package numDP__test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
600. 不含连续1的非负整数
https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/

给定一个正整数 n ，请你统计在 [0, n] 范围的非负整数中，有多少个整数的二进制表示中不存在 连续的 1 。
输入: n = 5
输出: 5
解释:
下面列出范围在 [0, 5] 的非负整数与其对应的二进制表示：
0 : 0
1 : 1
2 : 10
3 : 11
4 : 100
5 : 101
其中，只有整数 3 违反规则（有两个连续的 1 ），其他 5 个满足规则。

方法：数位dp枚举
1.5可以写成101。
2.向之前的数位dp一样，从高位向低位取（可以用高位的值压制低位的取值范围）高位在左，低位在右。
3.所有的数字都可以写成长度为3的字符串（以[0...5]为例）

hasLimit的作用是为了在枚举的过程中保证数字小于等于数字可以达到的最大上界。
preIsOne保证不会让两个1相邻。
*/
func findIntegers(n int) int {

	s := fmt.Sprintf("%b", n)

	var numDP func(i int, hasLimit bool, preIsOne bool, num string) (res int)
	numDP = func(i int, hasLimit bool, preIsOne bool, num string) (res int) {
		if i == len(s) {
			fmt.Println(num)
			return 1
		}

		if hasLimit {
			// 如果选0
			res += numDP(i+1, s[i] == '0', false, num+strconv.Itoa(0))
			// 如果选1
			if s[i] == '1' && preIsOne == false {
				res += numDP(i+1, true, true, num+strconv.Itoa(1))
			}

		} else {
			// 如果选0
			res += numDP(i+1, false, false, num+strconv.Itoa(0))
			// 如果选1
			if preIsOne == false {
				res += numDP(i+1, false, true, num+strconv.Itoa(1))
			}
		}
		return res
	}
	return numDP(0, true, false, "")

}

/*
记忆化
*/
func findIntegers2(n int) int {

	s := fmt.Sprintf("%b", n)
	// mem 存储 i、preIsOne的状态 hasLimit只管false的状态
	mem := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		mem[i] = []int{-1, -1}
	}

	var numDP func(i int, hasLimit bool, preIsOne int) (res int)
	numDP = func(i int, hasLimit bool, preIsOne int) (res int) {
		if i == len(s) {
			return 1
		}
		if hasLimit == false {
			if mem[i][preIsOne] != -1 {
				return mem[i][preIsOne]
			}
			defer func() {
				mem[i][preIsOne] = res
			}()
		}

		if hasLimit {
			// 如果选0
			res += numDP(i+1, s[i] == '0', 0)
			// 如果选1
			if s[i] == '1' && preIsOne == 0 {
				res += numDP(i+1, true, 1)
			}

		} else {
			// 如果选0
			res += numDP(i+1, false, 0)
			// 如果选1
			if preIsOne == 0 {
				res += numDP(i+1, false, 1)
			}
		}
		return res
	}
	return numDP(0, true, 0)

}

func TestFindInteger(t *testing.T) {
	fmt.Println(findIntegers2(5)) // 101
	//fmt.Println(findIntegers(2))
}
