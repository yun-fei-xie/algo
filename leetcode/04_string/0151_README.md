# 151. 反转字符串中的单词

https://leetcode.cn/problems/reverse-words-in-a-string/description/


## 解题思路

1. 拆分字符串中的单词 
2. 以单词作为反转的基本单位

单词与单词之间可能会有多个空白符
单纯的split不可行（使用正则?）

这里使用strings提供的Fields函数。

## 解题代码

```go
func reverseWords(s string) string {

  	sb := strings.Builder{}
	splitStrings := strings.Fields(s)
	for i := len(splitStrings) - 1; i >= 0; i-- {
		sb.WriteString(splitStrings[i])
		if i != 0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}
```

