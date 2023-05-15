package subset

import (
	"fmt"
	"testing"
)

/*
给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。

返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。



示例 1：
输入：s = "a1b2"
输出：["a1b2", "a1B2", "A1b2", "A1B2"]
示例 2:

输入: s = "3z4"
输出: ["3z4","3Z4"]


提示:
1 <= s.length <= 12
s 由小写英文字母、大写英文字母和数字组成

方法：回溯
对于每个字符，有如下情况
1. 不是字母->(原封不动、转换大小写)
2. 是字母->(原封不动、转换大小写)

于是可以得到，对于每个字母
1. 不做处理，加入路径中。进入下一轮递归。
2. 回溯回来之后，如果是字母，进行大小写互换。


*/

func letterCasePermutation(s string) []string {
	var ans = make([]string, 0)
	var path = make([]uint8, 0)
	var length = len(s)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if startIndex == length {
			temp := make([]uint8, len(path))
			copy(temp, path)
			ss := string(temp)
			ans = append(ans, ss)
			return
		}
		// 当前字符无需处理，直接加入路径
		path = append(path, s[startIndex])
		dfs(startIndex + 1)
		path = path[0 : len(path)-1]

		// 如果当前字符是小写
		if isAphaLower(s[startIndex]) {
			path = append(path, s[startIndex]-'a'+'A')
			dfs(startIndex + 1)
			path = path[0 : len(path)-1]

		} else if isAphaUpper(s[startIndex]) {
			// 如果当前字符是大写
			path = append(path, s[startIndex]-'A'+'a')
			dfs(startIndex + 1)
			path = path[0 : len(path)-1]
		}

	}
	dfs(0)
	return ans

}

func isAphaLower(ch uint8) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	return false
}

func isAphaUpper(ch uint8) bool {
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}

func TestLetterCasePermutation(t *testing.T) {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4B"))
}
