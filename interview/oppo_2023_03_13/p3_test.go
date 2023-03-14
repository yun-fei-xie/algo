package oppo_2023_03_13

import (
	"fmt"
	"testing"
)

func tourismRoutePlanning(scenicspot int) int64 {
	// write code here

	mem := make([]int64, scenicspot)

	var search func(n int) int64
	search = func(n int) int64 {
		if n == 1 || n == 2 {
			return int64(n)
		}

		if mem[n-1] == 0 {
			mem[n-1] = search(n - 1)
		}

		if mem[n-2] == 0 {
			mem[n-2] = search(n - 2)
		}

		return mem[n-1] + mem[n-2]
	}

	return search(scenicspot)
}

func TestTourismRoutePlanning(t *testing.T) {

	fmt.Println(tourismRoutePlanning(4))

}
