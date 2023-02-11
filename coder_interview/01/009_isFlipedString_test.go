package _1

import (
	"fmt"
	"strings"
	"testing"
)

/**
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

示例1:

 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True
示例2:

 输入：s1 = "aa", s2 = "aba"
 输出：False
提示：

字符串长度在[0, 100000]范围内。
说明:

你能只调用一次检查子串的方法吗？

"waterbottlewaterbottle"
"erbottlewat"


*/

func flipedAux(s string, index int) string {
	// 让下标转起来  i%len(s)  index[0 , 6]
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		sb.WriteByte(s[(index+i)%len(s)])
	}
	return sb.String()
}

func TestFlipedAux(t *testing.T) {
	s := "abcdefg"
	fmt.Println(flipedAux(s, 2))

}

func isFlipedString1(s1 string, s2 string) bool {
	if s1 == "" && s2 == "" || (s1 == s2) {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	for index := 0; index < len(s1); index++ {
		fliString := flipedAux(s1, index)
		if fliString == s2 {
			return true
		}
	}
	return false
}

// 更加高级的解法
/**
s1+s1得到的拼接字符串包含了 s1所有旋转可能包含的子串
因此，只需要判断s2是否是字符串s1+s1的子串即可
*/

func isFlipedString2(s1 string, s2 string) bool {

	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)

}

func TestIsFlipedString(t *testing.T) {
	s1 := "waterbottle"
	s2 := "erbottlewat"

	res := isFlipedString1(s1, s2)
	fmt.Println(res)

	res2 := isFlipedString2(s1, s2)
	fmt.Println(res2)

}
