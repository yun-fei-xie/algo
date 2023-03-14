package easy

import (
	"fmt"
	"math/bits"
	"testing"
)

/*
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
示例 1：
输入：x = 1, y = 4
输出：2
解释：
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
上面的箭头指出了对应二进制位不同的位置。

解题思路：先对两个数字进行异或运算，然后统计计算结果中二进制数1个个数。
这里重点是如何统计1的个数。

1. 内置函数
2. 造轮子
	位运算，和1进行与运算可以拿到最后一位是0还是1。向右移动一位去掉末位可以使用>>
3. Brian Kernighan 算法

我们需要循环右移 888 次才能得到答案。而实际上如果我们可以跳过两个 111 之间的 000，直接对 111 进行计数，那么就只需要循环 333 次即可。
我们可以使用 Brian Kernighan\text{Brian Kernighan}Brian Kernighan 算法进行优化，具体地，该算法可以被描述为这样一个结论：记 f(x)f(x)f(x) 表示 xxx 和 x−1x-1x−1 进行与运算所得的结果（即 f(x)=x & (x−1)f(x)=x~\&~(x-1)f(x)=x & (x−1)），那么 f(x)f(x)f(x) 恰为 xxx 删去其二进制表示中最右侧的 111 的结果。
基于该算法，当我们计算出 s=x⊕ys = x \oplus ys=x⊕y，只需要不断让 s=f(s)s = f(s)s=f(s)，直到 s=0s=0s=0 即可。
这样每循环一次，sss 都会删去其二进制表示中最右侧的 111，最终**循环的次数**即为 sss 的二进制表示中 111 的数量。

*/

// 只用内置的函数
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func hammingDistance2(x int, y int) int {
	ans := 0
	for s := x ^ y; s > 0; s = s >> 1 {
		ans += s & 1
	}
	return ans
}

func hammingDistance3(x int, y int) int {
	ans := 0
	s := x ^ y
	for s > 0 { // 每次f(x) = x & (x-1) 会消除最右侧的bit 1 , 因此当f(x) > 0 时，循环的次数就是bit 1的个数。
		s = s & (s - 1)
		ans++
	}
	return ans
}

func TestHammingDist(t *testing.T) {
	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance2(1, 4))
	fmt.Println(hammingDistance3(1, 4))
}
