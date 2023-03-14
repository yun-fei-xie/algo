package _36

import (
	"fmt"
	"testing"
)

func vowelStrings(words []string, left int, right int) int {
	var res = 0
	for i := left; i <= right; i++ {

		word := words[i]

		first := word[0]
		last := word[len(word)-1]
		if (first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u') && (last == 'a' || last == 'e' || last == 'i' || last == 'o' || last == 'u') {
			res++
		}
	}

	return res
}

func TestVowelString(t *testing.T) {
	fmt.Println(vowelStrings([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4))
}
