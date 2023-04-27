package _4_string

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/longest-common-prefix/

输入：strs = ["flower","flow","flight"]
输出："fl"

纵向扫描
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ { //以第一个字符串为基准
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || strs[j][i] != strs[0][i] { //前半部分条件是长度、后半部分条件是字符对比
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
