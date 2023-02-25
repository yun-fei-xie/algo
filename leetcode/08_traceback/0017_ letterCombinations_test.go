package _8_traceback

import (
	"fmt"
	"strconv"
	"testing"
)

/*
https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
*/

var digitsMap = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var res = []string{}
	var path string = ""
	var ss = make([]string, 0)

	// 拿到需要被遍历的字符串
	for i := 0; i < len(digits); i++ {
		char := digits[i]
		index, _ := strconv.Atoi(string(char))
		ss = append(ss, digitsMap[index])
	}
	// 遍历
	var dfs func(depth int)
	dfs = func(depth int) {
		if depth == len(digits) {
			res = append(res, path)
			return
		}

		s := ss[depth]

		for i := 0; i < len(s); i++ { // 回溯的核心
			path = path + string(s[i])
			dfs(depth + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}

func TestLetterCombinations(t *testing.T) {
	digits := ""
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
