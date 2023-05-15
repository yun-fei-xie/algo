package subset

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
*/

var numberToDigits = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}
	var path []uint8
	var maxDepth = len(digits)
	var traceback func(startIndex int, depth int)
	traceback = func(startIndex int, depth int) {
		if depth == maxDepth {
			s := string(path)
			ans = append(ans, s)
			return
		}
		digitIndex := digits[startIndex] - '0'      //拿到数字
		numberToDigit := numberToDigits[digitIndex] // 拿到字符串
		for i, l := 0, len(numberToDigit); i < l; i++ {
			path = append(path, numberToDigit[i])
			traceback(startIndex+1, depth+1)
			path = path[:len(path)-1]
		}

	}
	traceback(0, 0)
	return ans
}

func TestLetterCombinations(t *testing.T) {
	digits := "23"
	res := letterCombinations(digits)
	fmt.Println(res)

}

func TestChar(t *testing.T) {
	//s := "hello"
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(string(s[i]))
	//}
	s := ""
	fmt.Println(len(s))

}
