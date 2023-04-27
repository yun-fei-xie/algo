package _4_string

import (
	"fmt"
	"strings"
	"testing"
)

/*
对于溢出的处理方式通常可以转换为 INT_MAX 的逆操作。比如判断某数乘以 101010 是否会溢出，那么就把该数和 INT_MAX 除以 101010 进行比较。
*/
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	return -1
}

func TestMyAtoi(t *testing.T) {

	s := "    hello"
	fmt.Println(strings.TrimLeft(s, " "))
}
