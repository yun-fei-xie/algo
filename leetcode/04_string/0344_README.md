# 解题思路
左右双指针left ,right 
不断交换left和right指向的字符的位置以达到reverse的目的。



## 解题代码

```go
func reverseString(s []byte) {

	left := 0
	right := len(s) - 1

	for left < right {

		s[left], s[right] = s[right], s[left] // golang的语法糖
		left++
		right--
	}

}
```