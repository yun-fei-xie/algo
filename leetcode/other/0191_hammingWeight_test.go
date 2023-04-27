package other

import (
	"fmt"
	"testing"
)

func hammingWeight(num uint32) int {
	ans := 0

	for num != 0 {
		if num&1 == 1 {
			ans++
		}
		num = num >> 1
	}
	return ans
}

func TestHammingWeight(t *testing.T) {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
	fmt.Println(hammingWeight(00000000000000000000000010000000))
}
