# 383. 赎金信

## 题目链接

https://leetcode.cn/problems/ransom-note/

## 题目描述

给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。

## 解题思路

属于查找问题，

## 解题代码


```go
func canConstruct(ransomNote string, magazine string) bool {
	m1 := [26]int{}
	m2 := [26]int{}

	for _, char1 := range ransomNote {
		m1[char1-'a']++
	}
	for _, char2 := range magazine {
		m2[char2-'a']++
	}
	for i := 0; i < len(m1); i++ {
		if m2[i] < m1[i] {
			return false
		}
	}
	return true
}
```

