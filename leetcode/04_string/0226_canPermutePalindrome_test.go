package _4_string

/*
如果一个字符串可以组成一个回文串，那么：
(1) 如果它的长度为偶数，那么每个字符都必须出现偶数次；
(2) 如果它的长度为奇数，那么除了一个字符出现奇数次以外，其它的字符都必须出现偶数次。
因此可以总结得到，如果一个字符串可以组成一个回文串，那么出现奇数次的字符的数量不能超过 1。
*/
func canPermutePalindrome(s string) bool {
	var cnts = [256]int{}
	for _, char := range s {
		cnts[char-'0']++
	}
	var flag = false
	for _, num := range cnts {
		if num%2 != 0 {
			if flag {
				return false
			}
			flag = true
		}
	}
	return true
}
