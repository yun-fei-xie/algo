package _105_double_test

import (
	"fmt"
	"math"
	"testing"
)

func minExtraChar(s string, dictionary []string) int {
	mp := make(map[string]struct{})
	for _, word := range dictionary {
		mp[word] = struct{}{}
	}
	// 记忆化
	mem := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		mem[i] = -1
	}

	// 从startIndex开始向后切分能够拿到多少长度的字符串
	var dfs func(startIndex int) int
	dfs = func(startIndex int) int {
		if startIndex == len(s) {
			return 0
		}
		if mem[startIndex] != -1 {
			return mem[startIndex]
		}
		maxLens := math.MinInt
		defer func() {
			mem[startIndex] = maxLens
		}()

		// i是分割点，将s[startIndex:len(s)]分割成两个部分 s [startIndex...i] 和[i+1...len(s)-1] 两个部分
		for i := startIndex; i < len(s); i++ {
			// 在里面
			if _, found := mp[s[startIndex:i+1]]; found {
				maxLens = max(maxLens, (i-startIndex+1)+dfs(i+1))
			} else {
				maxLens = max(maxLens, dfs(i+1))
			}
		}
		return maxLens
	}
	maxLen := dfs(0)
	return len(s) - maxLen
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

func TestMinExtraChar(t *testing.T) {
	fmt.Println(minExtraChar("sayhelloworld", []string{"hello", "world"}))
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	fmt.Println(minExtraChar("voctvochpgutoywpnafylzelqsnzsbandjcqdciyoefi",
		[]string{"tf", "v", "wadrya", "a", "cqdci", "uqfg", "voc", "zelqsn", "band", "b", "yoefi", "utoywp", "herqqn", "umra", "frfuyj", "vczatj", "sdww"}))
}
