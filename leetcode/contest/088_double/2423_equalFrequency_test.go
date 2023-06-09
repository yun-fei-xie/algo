package _88_double

import (
	"fmt"
	"sort"
	"testing"
)

/*
2423. 删除字符使频率相同
https://leetcode.cn/problems/remove-letter-to-equalize-frequency/description/

给你一个下标从 0 开始的字符串 word ，字符串只包含小写英文字母。你需要选择 一个 下标并 删除 下标处的字符，使得 word 中剩余每个字母出现 频率 相同。
如果删除一个字母后，word 中剩余所有字母的出现频率都相同，那么返回 true ，否则返回 false 。
注意：
字母 x 的 频率 是这个字母在字符串中出现的次数。
你 必须 恰好删除一个字母，不能一个字母都不删除。

方法：分类讨论
这是一套披着外衣的不简单的简单题


*/

type pair struct {
	freq int
	chas int
}

func equalFrequency(word string) bool {

	hashMap := make(map[int32]int)
	for _, char := range word {
		hashMap[char]++
	}
	//1.全场只有一种字符 2.全场每个字符都只出现了一次
	if len(hashMap) == 1 || len(hashMap) == len(word) {
		return true
	}

	//  [a->23 , b->23 , c->22 , d->22]
	mp := make(map[int]int)
	for _, freq := range hashMap {
		mp[freq]++
	}

	// 最后必然只剩下两种频率的字符
	if len(mp) != 2 {
		return false
	}
	ps := make([]pair, 0)
	for key, val := range mp {
		ps = append(ps, pair{
			freq: key,
			chas: val,
		})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].freq < ps[j].freq
	})

	// 频率高的只有一个字符，并且这个字符的频率比其他字符高1
	if ps[1].freq == ps[0].freq+1 && ps[1].chas == 1 {
		return true
	}

	// 频率低的只有一个字符，并且这个字符的频率只有1
	if ps[0].freq == 1 && ps[0].chas == 1 {
		return true
	}

	return false
}

func TestEqualFrequency(t *testing.T) {
	//fmt.Println(equalFrequency("abcc"))
	//fmt.Println(equalFrequency("aazz"))
	//fmt.Println(equalFrequency("zzzz"))
	//fmt.Println(equalFrequency("abcd"))
	fmt.Println(equalFrequency("abbcc"))
}
