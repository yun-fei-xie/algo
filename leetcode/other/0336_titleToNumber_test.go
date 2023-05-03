package other

import (
	"fmt"
	"testing"
)

/*
本质是进制转换
*/

func titleToNumber(columnTitle string) int {
	var ans int
	for i := 0; i < len(columnTitle); i++ {
		ans = ans*26 + (int(columnTitle[i]-'A') + 1)
	}
	return ans
}

func TestTitleToNumber(t *testing.T) {
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("AC"))
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("B"))
}
