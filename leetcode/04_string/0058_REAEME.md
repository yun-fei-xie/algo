# 剑指 Offer 58 - II. 左旋转字符串

https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/


## 解题思路

本质就是一个数组的旋转。
如果长度为6 左旋转k 相当于右旋转 length-k  而k需要做处理 k = k %length
例如，当length=6 旋转6位和旋转12位并没有区别。

于是就有 k = k %length   m = length - k  向右旋转m
如果之前的位置位i 新的位置位 (i+m）%（length）
注意边界，如果length = 7 ,i=6 向右1位-> 应该到0
（6+1）% 7 = 0



## 解题代码

这里使用额外的空间做这件事情。

如果不使用额外的空间：
局部翻转+整体翻转
1. 反转区间为前n的子串
2. 反转区间为n到末尾的子串
3. 反转整个字符串



```go

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	res := make([]byte, len(b))
	n = n % len(b)
	m := len(b) - n

	for i := 0; i < len(b); i++ {

		res[(i+m)%len(b)] = b[i]
	}
	return string(res)
}

```