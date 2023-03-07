# leetcode 202快乐数

## 题目链接

https://leetcode.cn/problems/happy-number/

## 题目描述

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

## 解题思路

1. 模拟题目描述的过程，看看最后的结果是不是1。
2. 如何处理无限循环呢？用一个hashtable 记录每次计算的结果。如果当前计算的结果曾经出现过，那么这个计算过程将会无限循环。

## 解题代码

```go
func isHappy(n int) bool {
	record := make(map[int]struct{})
	square := n
	record[square] = struct{}{}
	for square != 1 {
		square = Aux(square)
		if _, found := record[square]; found {
			return false
		}

		record[square] = struct{}{}
	}
	return true
}

func Aux(n int) int {
	sum := 0
	for n != 0 {
		mod := n % 10
		n = n / 10
		sum += mod * mod
	}
	return sum
}
```
