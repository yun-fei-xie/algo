package _2_slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

/*
输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], minutes = 3
输出：16
解释：书店老板在最后 3 分钟保持冷静。
感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.
*/

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	sum := 0
	for i := 0; i < len(customers); i++ { // 老板不生气带来的客人数量
		if grumpy[i] == 0 {
			sum += customers[i]
		}
	}
	// 窗口滑动
	max := math.MinInt64
	left := 0
	right := 0
	sum2 := 0
	// 固定窗口
	for right < len(customers) {
		if right < minutes { // init
			sum2 += customers[right] * grumpy[right] // 不生气是0 生气是1（才会被累加）
			right++
		} else { // 第一次进入 right == minutes 长度比窗口大1
			if sum2 > max {
				max = sum2
			}
			sum2 += customers[right] * grumpy[right]
			sum2 -= customers[left] * grumpy[left]
			left++
			right++
		}
	}
	return sum + max
}

func TestMaxSatisfied(t *testing.T) {
	res := maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)
	fmt.Println(res)
}
