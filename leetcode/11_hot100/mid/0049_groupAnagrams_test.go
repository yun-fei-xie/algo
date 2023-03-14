package mid

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/group-anagrams/?favorite=2cktkvj
49. 字母异位词分组
亚马逊
彭博 Bloomberg
微软 Microsoft
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

解题思路：当两个字符串包含的字母相同时，两个字符串互为字母异位词
可以使用这个特性，将排序后的字符串作为hash表的key。

这里要记住golang里面对单个字符串排序的方法(先转成byte数组，然后再进行排序)
*/

/*
方法一：排序
*/
func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string, 0)

	for i := 0; i < len(strs); i++ {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], strs[i])
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

/*
方法二：计数
为每个单词准备一个26位的数组，统计其每个字母出现的次数。互为异位词的数组元素必然相同。
因此，可以用数组作为同一组异位词的key。
*/
func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
}
