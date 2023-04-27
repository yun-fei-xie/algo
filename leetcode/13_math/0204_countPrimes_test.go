package _3_math

import (
	"fmt"
	"testing"
)

/*
暴力解法
在大于1的自然数中，除了1和它本身以外不再有其他因数的自然数。
因此对于每个数x，我们可以从小到大枚举[2,x−1]中的每个数y，判断y是否为x的因数。
*/
func countPrimes(n int) int {
	var valid func(m int) bool
	valid = func(m int) bool {
		for i := 2; i*i <= m; i++ { //[2 , 根号m]
			if m%i == 0 {
				return false
			}
		}
		return true
	}

	var ans int
	for i := 2; i < n; i++ {
		if valid(i) {
			ans++
		}
	}
	return ans
}

/*
埃氏筛,
[0...n) 题目说了，小于n。因此，不包含n。
*/

func countPrimes2(n int) int {
	prime := make([]bool, n)
	for i := 0; i < n; i++ {
		prime[i] = true
	}
	var ans int

	for i := 2; i < n; i++ {
		if prime[i] == true {
			ans++
			for j := 2 * i; j < n; j += i {
				prime[j] = false
			}
		}
	}
	return ans
}

func TestCountPrimes(t *testing.T) {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes2(10))
}
