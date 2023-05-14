package _45

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/neighboring-bitwise-xor/
6431. 相邻值的按位异或

下标从 0 开始、长度为 n 的数组 derived 是由同样长度为 n 的原始 二进制数组 original 通过计算相邻值的 按位异或（⊕）派生而来。

特别地，对于范围 [0, n - 1] 内的每个下标 i ：

如果 i = n - 1 ，那么 derived[i] = original[i] ⊕ original[0]
否则 derived[i] = original[i] ⊕ original[i + 1]
给你一个数组 derived ，请判断是否存在一个能够派生得到 derived 的 有效原始二进制数组 original 。

如果存在满足要求的原始二进制数组，返回 true ；否则，返回 false 。

二进制数组是仅由 0 和 1 组成的数组。

derived = [1,1,0]
似乎可以通过枚举+递推。
derived[i]=0 ,则原始的origin[i] 与 origin[i+1] 相同
derived[i]=1 ,则origin[i]与origin[i+1]不同
用回溯法，传入一个path数组

特殊情况最后一位派生有原始的第一位和最后一位通过异或得到

输入：derived = [1,1,0]
输出：true
解释：能够派生得到 [1,1,0] 的有效原始二进制数组是 [0,1,0] ：
derived[0] = original[0] ⊕ original[1] = 0 ⊕ 1 = 1
derived[1] = original[1] ⊕ original[2] = 1 ⊕ 0 = 1
derived[2] = original[2] ⊕ original[0] = 0 ⊕ 0 = 0


知道原始数组的第一个数，通过派生数组就可以推出原始数组的所有项。
枚举第一个位置是0还是1 是否满足条件


*/

func doesValidArrayExist(derived []int) bool {
	length := len(derived)
	ans := false
	path := make([]int, length)
	var traceback func(pos int)

	traceback = func(pos int) {
		// baseCase
		if pos == length-1 {
			if derived[pos] == path[pos]^path[0] {
				ans = true
			}
			return
		}

		preVal := path[pos]
		if derived[pos] == 0 {
			path[pos+1] = preVal
		} else {
			if preVal == 0 {
				path[pos+1] = 1
			} else {
				path[pos+1] = 0
			}
		}

		traceback(pos + 1)

	}

	traceback(0)
	return ans
}

func TestD6432(t *testing.T) {
	fmt.Println(doesValidArrayExist([]int{1, 1, 0}))
	fmt.Println(doesValidArrayExist([]int{1, 1}))
	fmt.Println(doesValidArrayExist([]int{1, 0}))
}
