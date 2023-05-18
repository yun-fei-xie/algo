package _4_string

import (
	"fmt"
	"strconv"
	"testing"
)

func addBinary(a string, b string) string {

	lena := len(a)
	lenb := len(b)
	maxLen := max(lenb, lena)

	a = reverseS(a)
	b = reverseS(b)

	var carry int
	var ans = ""
	for i := 0; i < maxLen; i++ {
		if i < lena {
			carry += int(a[i] - '0')
		}
		if i < lenb {
			carry += int(b[i] - '0')
		}
		ans += strconv.Itoa(carry % 2)
		carry = carry / 2

	}
	if carry != 0 {
		ans += "1"
	}
	return reverseS(ans)
}

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

func TestAddBinary(t *testing.T) {

	//fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("0", "0"))

}
