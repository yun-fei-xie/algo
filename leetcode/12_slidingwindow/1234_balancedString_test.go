package _2_slidingwindow

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/replace-the-substring-for-balanced-string/
有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
请返回待替换子串的最小可能长度。
如果原字符串自身就是一个平衡字符串，则返回 0。

输入：s = "QWER"
输出：0
解释：s 已经是平衡的了。

输入：s = "QQQW"
输出：2
解释：我们可以把前面的 "QQ" 替换成 "ER"。
*/

/*
字符串的长度已经知道了，所以每个字符应该出现 m= len(s)/4次
如果子串s[left ,right] 之外的每个字符出现的次数都<=m ,
那么一定可以通过替换当前子串s[left,right]使得s成为一个平衡字符串。（具体怎么替换不管）
也就是说,如果s[left,right]之外的每个字符出现的次数都<=m ,那么 s[left,right]就是一个候选解。

这个题和1658都是一个反向思维的题目。
原始问题不容易解决，寻找等价的反向问题。
*/
func balancedString(s string) int {
	m := len(s) / 4
	counter := ['X']int{}
	for _, ch := range s {
		counter[ch]++
	}
	if counter['Q'] == m && counter['W'] == m && counter['E'] == m && counter['R'] == m {
		return 0
	}

	// 滑动窗口
	var ans int = len(s)
	for left, right := 0, 0; right < m*4; right++ {
		// counter记录窗口之外的字符出现的次数 right进窗口，窗口之外--
		counter[s[right]]--
		for counter['Q'] <= m && counter['W'] <= m && counter['E'] <= m && counter['R'] <= m {
			ans = min(ans, right-left+1)
			counter[s[left]]++
			left++
		}
	}
	return ans
}

func TestBalancedString(t *testing.T) {
	fmt.Println(balancedString("QQWE"))
	fmt.Println(balancedString("QQQW"))
}
