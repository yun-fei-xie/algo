package other

import (
	"testing"
)

/*
num!=0 这种写法有问题，不能处理前缀0 应该枚举每一位
输入n = 00000010100101000001111010011100
输出：964176192 (00111001011110000010100101000000)
如果不处理前导0，会导致结果的位数不对（因为num!=0）循环的次数不够。
*/

// 错误解法
func reverseBits(num uint32) uint32 {
	//	numStr := strconv.Itoa(int(num)) // golang没有内置的reverse
	var ans uint32 = 0
	for num != 0 {
		ans = (ans << 1) | (num & 1)
		num = num >> 1
	}

	return ans
}

// 正确解法
func reverseBits2(num uint32) uint32 {
	var ans uint32 = 0
	for i := 0; i < 32; i++ {
		ans = (ans << 1) | (num & 1) // nums&1拿到最后一位  ans<<1 向右移动一位给末尾留出  最后|运算将最后一位拼接上去
		num = num >> 1
	}
	return ans
}

func TestReverseBits(t *testing.T) {
	//fmt.Println(reverseBits(00000010100101000001111010011100))
	//fmt.Println(reverseBits(11111111111111111111111111111101))
}
