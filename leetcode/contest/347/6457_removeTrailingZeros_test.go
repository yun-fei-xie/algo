package _347

import (
	"fmt"
	"testing"
)

/*
51230100
*/
func removeTrailingZeros(num string) string {
	var i int
	for i = len(num) - 1; i >= 0; {
		if num[i] == '0' {
			i--
		} else {
			break
		}
	}
	return num[:i+1]
}

func TestRemoveTrailingZeros(t *testing.T) {
	fmt.Println(removeTrailingZeros("51230100"))
	fmt.Println(removeTrailingZeros("12345"))
	fmt.Println(removeTrailingZeros("0000"))
}
