# 剑指 Offer 05. 替换空格

https://leetcode.cn/problems/ti-huan-kong-ge-lcof/


## 解题思路

1. 识别空格
2. 替换成需要替换的字符



## 解题代码

```go

func replaceSpace(s string) string {
	b := []byte(s)
	sb := strings.Builder{}
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {  // 识别空格
			sb.WriteString("%20")
		} else {
			sb.WriteByte(b[i])
		}
	}
	return sb.String()
}



```