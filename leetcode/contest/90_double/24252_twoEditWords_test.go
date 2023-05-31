package _0_double_test

/*
2452. 距离字典两次编辑以内的单词
https://leetcode.cn/problems/words-within-two-edits-of-dictionary/description/

给你两个字符串数组 queries 和 dictionary 。数组中所有单词都只包含小写英文字母，且长度都相同。
一次 编辑 中，你可以从 queries 中选择一个单词，将任意一个字母修改成任何其他字母。从 queries 中找到所有满足以下条件的字符串：不超过 两次编辑内，字符串与 dictionary 中某个字符串相同。
请你返回 queries 中的单词列表，这些单词距离 dictionary 中的单词 编辑次数 不超过 两次 。单词返回的顺序需要与 queries 中原本顺序相同。

方法：二重循环暴力破解,字符串两两比较
*/
func twoEditWords(queries []string, dictionary []string) []string {
	var ans = make([]string, 0)
	for _, query := range queries {

		for j := 0; j < len(dictionary); j++ {
			if canReach(&query, &dictionary[j]) {
				ans = append(ans, query)
				break
			}
		}
	}
	return ans
}

func canReach(s1 *string, s2 *string) bool {
	length := len(*s1)
	diffCount := 0
	for i := 0; i < length; i++ {
		if (*s1)[i] != (*s2)[i] {
			diffCount++
		}
		if diffCount > 2 {
			return false
		}
	}
	return true
}
