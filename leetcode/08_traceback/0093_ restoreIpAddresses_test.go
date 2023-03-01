package _8_traceback

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/restore-ip-addresses/
递归四层（ip地址分为4段），把验证分割结果（中间辅以若干条件进行剪枝）
*/

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0)
	maxDepth := 4
	var dfs func(s string, startIndex int, depth int)
	dfs = func(s string, startIndex int, depth int) {
		if startIndex >= len(s) {
			if depth == maxDepth {
				sb := strings.Builder{}
				for i := 0; i < len(path); i++ {
					sb.WriteString(path[i])
					if i != len(path)-1 {
						sb.WriteString(".")
					}
				}
				res = append(res, sb.String())
			}
			return
		}

		for i := startIndex; i < len(s); i++ {
			if vaild(s[startIndex : i+1]) {
				path = append(path, s[startIndex:i+1])
				dfs(s, i+1, depth+1)
				path = path[0 : len(path)-1]
			}
		}
	}
	dfs(s, 0, 0)
	return res
}

func vaild(s string) bool {
	num, b := strconv.Atoi(s)
	if b != nil {
		return false
	}
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	return num >= 0 && num <= 255
}

func TestRestoreIp(t *testing.T) {
	s1 := "25525511135"
	res1 := restoreIpAddresses(s1)
	fmt.Println(res1)

	s2 := "0000"
	res2 := restoreIpAddresses(s2)
	fmt.Println(res2)

	s3 := "101023"
	res3 := restoreIpAddresses(s3)
	fmt.Println(res3)
}
