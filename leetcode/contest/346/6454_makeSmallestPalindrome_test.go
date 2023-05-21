package _46

import (
	"fmt"
	"testing"
)

func makeSmallestPalindrome(s string) string {
	byteString := []byte(s)
	for left, right := 0, len(byteString)-1; left <= right; {
		if byteString[left] > byteString[right] {
			byteString[left] = byteString[right]
		} else if byteString[left] < byteString[right] {
			byteString[right] = byteString[left]
		}
		left++
		right--
	}
	return string(byteString)
}

func TestMakeSmallestPalindrome(t *testing.T) {
	fmt.Println(makeSmallestPalindrome("egcfe"))
	fmt.Println(makeSmallestPalindrome("abcd"))
	fmt.Println(makeSmallestPalindrome("seven"))
}
