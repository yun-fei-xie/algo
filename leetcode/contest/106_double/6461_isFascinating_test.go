package _06_double

import (
	"fmt"
	"testing"
)

/*
优化点：哪些n一定不符合要求，<123  >329都是不合法的
*/
func isFascinating(n int) bool {
	if n == 0 {
		return false
	}
	twoN := 2 * n
	threeN := 3 * n

	visited := make([]int, 10) // 1-9

	for n != 0 {
		visited[n%10]++
		if visited[n%10] != 1 {
			return false
		}
		n = n / 10
	}

	for twoN != 0 {
		visited[twoN%10]++
		if visited[twoN%10] != 1 {
			return false
		}
		twoN = twoN / 10
	}

	for threeN != 0 {
		visited[threeN%10]++
		if visited[threeN%10] != 1 {
			return false
		}
		threeN = threeN / 10
	}

	if visited[0] != 0 {
		return false
	}
	return true
}

func TestIsFascinating(t *testing.T) {
	fmt.Println(isFascinating(192))
	fmt.Println(isFascinating(100))
}
