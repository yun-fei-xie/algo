package _348__test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
2719. 统计整数数目
https://leetcode.cn/problems/count-of-integers/

给你两个数字字符串 num1 和 num2 ，以及两个整数 max_sum 和 min_sum 。如果一个整数 x 满足以下条件，我们称它是一个好整数：
num1 <= x <= num2
min_sum <= digit_sum(x) <= max_sum.
请你返回好整数的数目。答案可能很大，请返回答案对 109 + 7 取余后的结果。
注意，digit_sum(x) 表示 x 各位数字之和。

方法：暴力解法，枚举从num1到num2的每一个数字，然后检测数位和是否满足要求，肯定会超时。
方法：构造+记忆化搜索（数位dp）

记忆化的时候，如果某种状态极度不平衡，可以省略这种状态。
*/
/*
构造+记忆化搜索
如何构造一个nums1到nums2这个范围内的数字，让它符合数位和在min_sum到max_sum之间
不妨从0开始构造，那么结果就是[0...nums2]这个范围内合法的数字 减去[0...nums1]这个范围内的数字。
整个问题转换为求[0...n]这个范围内有多少合法的数字。如何构造呢？
假设n="234"，尝试枚举每一位，第一位可以取{0,1,2}当第一位不取最高为2的时候，剩下的位数都可以枚举{0...9}。
在枚举的时候，可以使用一个bitSum记录从左边到当前累计的数位和。遇到不合法的结果直接return。
所以，这里的求解本质上是在枚举每一个位置上的可能的数组，求解出合法的组合。

哪里构成了重叠子问题？
举个例子：nums="45" ,min_sum=1,max_su=10
那么第一位可以选{0,1,2,3,4} 当第一位不是4（最高位）的时候，
第二位的可以选{0,1,2,...,8,9}。现在考虑{1,7}和{2,6}这两个组合。
假设他们还没有选齐所有的位数
1.那么两者后面需要选的位数相同
2.剩余可以使用的累加和也相同
3.并且二者在后面的位数的数字选择上都没有限制{0...9}
所以两者产生了重叠子问题。

*/

func count3(num1 string, num2 string, min_sum int, max_sum int) int {

	const mod int = 1e9 + 7
	var f func(s string) int
	f = func(s string) int {
		// 记忆化 两个s可能长度不同，因为需要分别构造mem
		mem := make([][]int, len(s))
		for i := 0; i < len(s); i++ {
			mem[i] = make([]int, max_sum+1)
			for j := 0; j <= max_sum; j++ {
				mem[i][j] = -1
			}
		}

		// 返回以i为起始位，符合条件的数位和的组合数量
		var traceBack func(i int, sum int, hasLimit bool) int
		traceBack = func(i int, sum int, hasLimit bool) int {
			if i == len(s) {
				if sum < min_sum || sum > max_sum {
					return 0
				}
				return 1
			}
			if sum > max_sum {
				return 0
			}

			var ret = 0
			if !hasLimit {
				if mem[i][sum] != -1 {
					return mem[i][sum]
				} else {
					defer func() {
						mem[i][sum] = ret
					}()
				}
			}

			if hasLimit {
				upper := int(s[i] - '0')
				for j := 0; j <= upper; j++ {
					if j == upper {
						ret += traceBack(i+1, sum+j, true) % mod
					} else {
						ret += traceBack(i+1, sum+j, false) % mod
					}
				}
			} else {
				for j := 0; j <= 9; j++ {
					ret += traceBack(i+1, sum+j, false) % mod
				}
			}
			return ret % mod
		}

		return traceBack(0, 0, true)
	}

	n1, _ := strconv.Atoi(num1)
	countN1 := 0
	for n1 != 0 {
		countN1 += n1 % 10
		n1 = n1 / 10
	}
	if countN1 >= min_sum && countN1 <= max_sum {

		return f(num2) - f(num1) + 1
	}
	return f(num2) - f(num1)
}

/*
不带记忆化的计算方式
*/
func count2(num1 string, num2 string, min_sum int, max_sum int) int {
	const mod int = 1e9 + 7
	var f func(s string) int
	f = func(s string) int {
		// 返回以i为起始位，符合条件的数位和的组合数量
		var traceBack func(i int, sum int, hasLimit bool) int
		traceBack = func(i int, sum int, hasLimit bool) int {
			if i == len(s) {
				if sum < min_sum || sum > max_sum {
					return 0
				}
				return 1
			}
			if sum > max_sum {
				return 0
			}

			var ret = 0
			if hasLimit {
				upper := int(s[i] - '0')
				for j := 0; j <= upper; j++ {
					if j == upper {
						ret += traceBack(i+1, sum+j, true)
					} else {
						ret += traceBack(i+1, sum+j, false)
					}
				}
			} else {
				for j := 0; j <= 9; j++ {
					ret += traceBack(i+1, sum+j, false)
				}
			}
			return ret % mod
		}

		return traceBack(0, 0, true)
	}

	n1, _ := strconv.Atoi(num1)
	countN1 := 0
	for n1 != 0 {
		countN1 += n1 % 10
		n1 = n1 / 10
	}
	if countN1 >= min_sum && countN1 <= max_sum {

		return f(num2) - f(num1) + 1
	}
	return f(num2) - f(num1)
}

/*
暴力枚举
*/
func count(num1 string, num2 string, min_sum int, max_sum int) int {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	var mod int = 10e9 + 7
	var countBit func(num int) int
	countBit = func(num int) int {
		var ans int
		for num != 0 {
			ans += num % 10
			num = num / 10
		}
		return ans
	}

	var hashMap = make(map[int]int)
	var ans int
	for i := n1; i <= n2; i++ {
		cb := countBit(i)
		if cb >= min_sum && cb <= max_sum {
			ans = (ans + 1) % mod
			hashMap[cb]++
		}
	}
	//fmt.Println(hashMap)
	return ans
}

func TestCount(t *testing.T) {
	fmt.Println(count("1", "9999", 1, 8))
	fmt.Println(count3("1", "9999", 1, 8))
	//fmt.Println(count2("1", "12", 1, 8))

}
