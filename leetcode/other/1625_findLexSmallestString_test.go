package other

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/lexicographically-smallest-string-after-applying-operations/description/

数字一旦超过9就会变成0 : 题干中的这句话有歧义。不是说累加超过9，直接置为0，然后重新开始。而是说9的下一位是0
比如当前是7，累加值是4 。累加完毕后结果其实是1(9 ,10->0 , 11->1)，而不是是0。

2023-03-19 改天再做。


暴力枚举
问题是如何枚举
题目给出s的长度为偶数

s = "5525", a = 9, b = 2
2555

如果b是偶数，那么无论轮转多少次，只能给s的奇数位做累加操作
如果b是奇数，那么可以给s的奇数和偶数做累加操作（就是每一位都可做累加）


从以上可以看出，做累加操作的次数和做轮转操作的次数是互相独立的，做轮转的次数并不会影响到能否对偶数位进行累加。
次数是独立的，但是结果并不是独立的。
因此我们可以先枚举做轮转的次数，然后再枚举做累加的次数，从而找到字典序最小的答案。

枚举做轮转的次数，然后令 t为 s 轮转后的字符串。
由于轮转最终会产生循环，且至多轮转 n 次（n 为 s 的长度），因此我们用一个数组 vis来记录每个位置是否被轮转过。
如果遇到之前轮转过的位置，则枚举结束。
对于每个 t，枚举对 t 的奇数位做累加操作的次数 j，再枚举对 t 的偶数位做累加操作的次数 k。
这里因为元素值范围是 [0,9][0,9][0,9]，因此我们做累加操作的次数上限是 999，再多势必会产生循环。
特殊的，如果b是偶数，则k的上限是0，否则是9。


abcabc

s1 ->

将s拼接到自己的后面可以实现快速的拿到轮转后的字符串。
怎么拿呢？想象一下，如果轮转一位，那么新的串的起始位其实是原始串的最后一位。
如果轮转2位，那么新的串的起始位其实是原始串的倒数第二位。
如果轮转k位，那么新的串的起始位其实是原始串的 k%len(s)


解题思路：
1. 轮转枚举和累加枚举是组合关系，需要用嵌套循环做
2. 外部循环先拿到轮转后的字符串,内部循环用一个for循环去依次对每一位需要累加位进行累加（累加的过程也是一个二重循环）
3. 如果轮转的位数b是奇数，那么每一位都会被轮转到。如果轮转的位数b是偶数，那么无论轮转多少次，都只能对s中的奇数位坐累加操作。
4. 也就是说，无论b是奇数还是偶数，轮转后的奇数位都会被累加到。所以可以先操作奇数位。如果b是奇数，再操作偶数位

这个题很容易调到坑里面就是，可以累加之后，基于累加的值做轮转。而不是基于一开始的值做轮转。




*/

func findLexSmallestString(s string, a int, b int) string {
	size := len(s)
	min := s
	visit := make([]bool, size)
	// 如何判断轮转出现循环 第一次轮转b,第一二次在原来的基础上再轮转b，此时相当于在最初的基础上轮转2b

	for i := 0; visit[i] == false; i = (i + b) % size { // i等于0就相当于原始的字符串
		visit[i] = true
		t := []byte(rotateString(s, i)) // 移动一次
		// 处理所有的奇数位
		for j := 0; j < 10; j++ {
			for p := 1; p < size; p += 2 {
				t[p] = '0' + (t[p]-'0'+uint8(j*a))%10 // 进行一次累加
			}
			if strings.Compare(min, string(t)) > 0 {
				min = string(t)
			}
		}
		// 处理所有的偶数位
		t = []byte(rotateString(s, i)) // 移动一次
		if b%2 == 1 {
			for j := 0; j < 10; j++ {
				for p := 0; p < size; p += 2 {
					t[p] = '0' + (t[p]-'0'+uint8(j*a))%10
				}
				if strings.Compare(min, string(t)) > 0 {
					min = string(t)
				}
			}
		}

	}
	return min
}

/*
将字符串s 向右移动k位
abcd-> abcdabcd
*/
func rotateString(s string, k int) string {
	l := len(s)
	k = k % l // 移动s.length 个长度后，字符串会回到原来的位置
	s = s + s

	// 移动k位之后，新的字符串的起始位置在哪里
	return s[l-k : (l - k + l)] // [l , r) 注意区间

}

func TestFindLexSmallestString(t *testing.T) {
	fmt.Println(findLexSmallestString("5525", 9, 2))
	fmt.Println(findLexSmallestString("74", 5, 1))
	fmt.Println(findLexSmallestString("0011", 4, 2))
	fmt.Println(findLexSmallestString("43987654", 7, 3))
}

func TestRotate(t *testing.T) {
	s := "1234"
	for i := 0; i < 10; i++ {
		sr := rotateString(s, 2)
		fmt.Println(sr)
		s = sr
	}

}
