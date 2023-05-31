package _0_double_test

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
2453. 摧毁一系列目标
https://leetcode.cn/problems/destroy-sequential-targets/

给你一个下标从 0 开始的数组 nums ，它包含若干正整数，表示数轴上你需要摧毁的目标所在的位置。同时给你一个整数 space 。
你有一台机器可以摧毁目标。给机器 输入 nums[i] ，这台机器会摧毁所有位置在 nums[i] + c * space 的目标，其中 c 是任意非负整数。你想摧毁 nums 中 尽可能多 的目标。
请你返回在摧毁数目最多的前提下，nums[i] 的 最小值 。

方法： 枚举每一个位置i，遍历nums[j]->[i<=j<=n-1],如果(nums[j]-nums[i])%space==0的话，nums[j]可以被nums[i]摧毁。
	这样的时间复杂度是O(n^2)级别的。会超时
方法：对nums的每一个数字按照space进行取模操作。相同模数的分为一组。例如，当space=10的时候，3、13、23、33...都可以被3摧毁。
	这些数字有一个特征，nums[i]%space =固定的数字。

*/

func destroyTargets1(nums []int, space int) int {
	sort.Ints(nums)
	var maxDestroy int
	var minNum int
	for i := 0; i < len(nums); i++ {
		//枚举每一个位置
		// if [nums[j]-nums[i]]%space ==0 可以被摧毁
		var count int
		for j := i; j < len(nums); j++ {
			if (nums[j]-nums[i])%space == 0 {
				count++
			}
		}
		if count > maxDestroy {
			maxDestroy = count
			minNum = nums[i]
		}
	}
	return minNum
}

func destroyTargets2(nums []int, space int) int {
	hashMap := make(map[int][]int, 0)

	for _, num := range nums {
		hashMap[num%space] = append(hashMap[num%space], num)
	}
	// 找出元素最多的那个slice里面的最小值
	maxLen, minNum := 0, math.MaxInt
	for _, arr := range hashMap {
		if len(arr) > maxLen {
			minNum = min(arr...)
			maxLen = len(arr)
		} else if len(arr) == maxLen {
			minNum = min(min(arr...), minNum)
		}
	}
	return minNum
}

func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

func TestDestroyTargets(t *testing.T) {
	//fmt.Println(destroyTargets2([]int{1, 3, 5, 2, 4, 6}, 2))
	//fmt.Println(destroyTargets2([]int{1, 5, 3, 2, 2}, 1000))
	fmt.Println(destroyTargets2([]int{3, 7, 8, 1, 1, 5}, 2))
}
