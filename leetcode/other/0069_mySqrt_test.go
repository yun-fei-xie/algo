package other

import (
	"fmt"
	"testing"
)

// 本题有多种解法：牛顿迭代、二分查找... 可以看题解 这里实现2种
// 第一种解法，思路最简单的逼近
func mySqrt(x int) int {
	var i int
	for i = 1; i*i <= x; i++ {

	}
	return i - 1
}

// 第二种解法：二分查找 从[1...X]中通过二分找出这个数字
// 相对于方法1，二分查找的速度要更加地快速
func mySqrt2(x int) int {

	var left = 1
	var right = x
	var ans int
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func TestMySqrt(t *testing.T) {
	//fmt.Println(mySqrt(10))
	//fmt.Println(mySqrt(9))
	//fmt.Println(mySqrt(16))
	//fmt.Println(mySqrt(22))
	fmt.Println(mySqrt2(10))
	fmt.Println(mySqrt2(9))
	fmt.Println(mySqrt2(16))
	fmt.Println(mySqrt2(22))
}
