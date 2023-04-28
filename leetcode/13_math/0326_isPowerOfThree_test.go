package _3_math

import (
	"fmt"
	"math"
	"testing"
)

/*
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
*/

/*
解法1， 循环除法。如果一个数字是3个幂。
那么，这个数字不断除以3，最后的结果一定是1。
例如，9/3->3  3/3->1
*/
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n = n / 3
	}
	return n == 1
}

/*
打表法，提前将3的幂存储起来
需要注意的是，1也是3的幂。3^0=1
*/
func isPowerOfThree2(n int) bool {
	set := make(map[int]struct{})
	set[1] = struct{}{}
	for i := 3; i < math.MaxInt64/3; i = i * 3 {
		set[i] = struct{}{}
	}

	if _, ok := set[n]; ok {
		return true
	}
	return false
}

/*
递归写法
*/
func isPowerOfThree3(n int) bool {
	if n < 1 {
		return false
	}
	var rec func(n int) bool
	rec = func(n int) bool {
		if n == 1 {
			return true
		}
		if n%3 != 0 {
			return false
		}
		return rec(n / 3)
	}
	return rec(n)
}
func TestIsPowerOfThree(t *testing.T) {
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(1))
}
