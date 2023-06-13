package _49

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/contest/weekly-contest-349/problems/collecting-chocolates/

先试试贪心
为了获取这个类型的巧克力，尝试所有的位置，取总成本最低的那个位置.
*/
func minCost(nums []int, x int) int64 {
	// 找出最小成本的类型所在的位置
	length := len(nums)
	minIndex := 0
	minNum := nums[0]
	for i := 1; i < length; i++ {
		if nums[i] < minNum {
			minIndex = i
		}
	}
	var ans int
	ans += nums[minIndex] // 先拿下最小成本类型的巧克力
	//
	rotateCnt := 0
	for d := 1; d <= length-1; d++ {
		currentIndex := ((minIndex - d) + length) % length
		// 根据之前旋转的次数找出当前真实的位置
		realIndex := (currentIndex + rotateCnt + length) % length
		// for循环 找出最小成本 最多转到minIndex ，因为minIndex后面的成本都比minIndex大
		// k表示当前位置最多可以旋转多少次
		var spanMax = 0
		if realIndex < minIndex {
			spanMax = minIndex - realIndex
		} else {
			spanMax = realIndex - minIndex
		}
		mCost := math.MaxInt
		spanCnt := 0
		// 取当前成本最小的
		for k := 0; k < spanMax; k++ {
			if nums[(realIndex+k+length)%length]+k*x < mCost {
				mCost = nums[(realIndex+k+length)%length] + k*x
				spanCnt = k
			}
		}
		ans += mCost
		rotateCnt += spanCnt
	}
	return int64(ans)
}

func TestMinCost(t *testing.T) {
	fmt.Println(minCost([]int{20, 1, 15}, 5))
}
