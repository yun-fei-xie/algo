package _1

/*
URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）

示例 1：
输入："Mr John Smith    ", 13
输出："Mr%20John%20Smith"

示例 2：
输入："               ", 5
输出："%20%20%20%20%20"

*/

import (
	"fmt"
	"strings"
	"testing"
)

func replaceSpaces(S string, length int) string {
	temp := S[0:length]
	return strings.Replace(temp, " ", "%20", -1)
}

func TestReplaceSpace(t *testing.T) {

	fmt.Println(replaceSpaces("               ", 5))
	fmt.Println(replaceSpaces("Mr John Smith    ", 13))

}
