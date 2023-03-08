# 541. 反转字符串 II

https://leetcode.cn/problems/reverse-string-ii/description/


## 解题思路

题目要求每计数2k个字符，就翻转前k个字符。
最容易想到的思路就是模拟。
如何模拟呢？因为有时候并没有2k个字符。
设置规律：翻转->不翻转->翻转->不翻转... 以此往复。

## 解题代码

```go
func reverseStr(s string, k int) string {
reverse := true // 翻转->不翻转->翻转...
sb := strings.Builder{}
var i int = 0
for i = 0; i < len(s); i = i + k {
if reverse {
reverseSubString := reverseS(s[i:min(i+k, len(s))]) // [left , right)
sb.WriteString(reverseSubString)
reverse = false
} else {
sb.WriteString(s[i:min(i+k, len(s))])
reverse = true
}
}

//sb.WriteString(s[i-k : min(i, len(s))])

return sb.String()
}

// 这里把参数修改为[]byte 就可以使用双指针的方式翻转传入的字符串
func reverseS(s string) string {
sb := strings.Builder{}
for i := len(s) - 1; i >= 0; i-- {
sb.WriteByte(s[i])
}
return sb.String()
}

func min(i, j int) int {
if i > j {
return j
} else {
return i
}

}
```