package _3_math

import (
	"fmt"
	"math"
	"testing"
)

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n%2 == 0 {
		n = n / 2
	}
	return n == 1
}

func isPowerOfTwo2(n int) bool {
	set := make(map[int]struct{})
	set[1] = struct{}{}
	for i := 2; i < math.MaxInt64/2; i = i * 2 {
		set[i] = struct{}{}
	}

	if _, ok := set[n]; ok {
		return true
	}
	return false
}
func TestIsPowerOfTwo(t *testing.T) {
	fmt.Println(isPowerOfTwo2(2))
	fmt.Println(isPowerOfTwo2(3))
	fmt.Println(isPowerOfTwo2(8))
	fmt.Println(isPowerOfTwo2(12))
}
