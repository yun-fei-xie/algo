package _0_dp

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
示例 1：
输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。

思路:最容易想到的暴力解法-> 双重循环
*/
func findLength(nums1 []int, nums2 []int) int {

	var maxLength int = 0
	l1, l2 := len(nums1), len(nums2)
	for i := 0; i < l1; i++ {

		for j := 0; j < l2; j++ {
			// 第一次匹配上之后，双指针向后迭代
			if nums1[i] == nums2[j] {
				var commLen = 1
				for n1, n2 := i+1, j+1; n1 < l1 && n2 < l2 && nums1[n1] == nums2[n2]; {
					commLen++
					n1++
					n2++
				}
				maxLength = max(maxLength, commLen)
			}
		}
	}
	return maxLength
}

/*
使用递归解决
*/

/*
使用动态规划解决问题
dp[i][j] ：长度为i，末尾项为A[i-1]的子数组，与长度为j，末尾项为B[j-1]的子数组，二者的最大公共后缀子数组长度。
如何初始化数组？
   1  2  3  2  1  nums1
3  0  0  1  0  0
2  0  1  0  2  0
1  1  0  1  0  3
4  0  1  0  1  0
7  0  0  1  0  1
nums2

for i :=0 ;i< m ;i++ {
	for j:=0 ;j<n ;j++ {
		dp[i][j] = dp[i-1][j-1] + {0 <- nums1[i] != nums2[j] , 1 <- nums1[i]==nums2[j] }
  }
}
在代码中对第一行和第一列进行初始化
dp[i][j] = dp[i-1][j-1] + {0 <- nums1[i] != nums2[j] , 1 <- nums1[i]==nums2[j]
这行代码就相当于暴力解法中的第三重循环中的双指针

*/

func findLengthDp(nums1 []int, nums2 []int) int {
	lenNum1, lenNum2 := len(nums1), len(nums2)
	dp := make([][]int, lenNum2)
	for i := 0; i < lenNum2; i++ {
		dp[i] = make([]int, lenNum1)
	}

	var maxLen = 0
	// 初始化第一行
	for i := 0; i < lenNum1; i++ {
		if nums1[i] == nums2[0] {
			dp[0][i] = 1
		}
		maxLen = max(maxLen, dp[0][i])
	}
	// 初始化第一列
	for j := 0; j < lenNum2; j++ {
		if nums2[j] == nums1[0] {
			dp[j][0] = 1
		}
		maxLen = max(maxLen, dp[j][0])
	}
	// 递推其他位置
	for i := 1; i < lenNum2; i++ {
		for j := 1; j < lenNum1; j++ {
			if nums2[i] == nums1[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = 0
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}

	return maxLen
}

func TestFindLength(t *testing.T) {
	//fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	fmt.Println(findLengthDp([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	//fmt.Println(findLength([]int{0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))
	//fmt.Println(findLengthDp([]int{0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))
	//fmt.Println(findLengthDp([]int{1, 2, 3, 2, 8}, []int{5, 6, 1, 4, 7}))

}
