package other

import (
	"fmt"
	"testing"
)

/*
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
*/
func convertToTitle(columnNumber int) string {
	var ans = ""
	for {
		columnNumber = columnNumber - 1
		ans = string('A'+uint8(columnNumber%26)) + ans
		columnNumber = columnNumber / 26
		if columnNumber == 0 {
			break
		}
	}
	return ans
}

func TestConvertToTitle(t *testing.T) {
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(33))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(2147483647))
}
