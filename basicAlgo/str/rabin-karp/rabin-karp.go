package rabin_karp

/*
RabinKarp 假设原串和模式串都是由26个小写字母构成
*/
func RabinKarp(text string, pattern string) int {

	m := len(text)
	n := len(pattern)
	patternHash := hashEncode(pattern)
	for i := 0; i <= (m - n); i++ {
		//text中的每个子串为：text[i:n+i]
		if hashEncode(text[i:n+i]) == patternHash {
			return i
		}
	}

	return -1
}

/*
hashEncode [a...z]->[0...25]编码
*/
func hashEncode(s string) int {
	var h int
	for i := 0; i < len(s); i++ {
		h = h*26 + int(s[i]-'a')
	}
	return h
}
