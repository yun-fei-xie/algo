package _6_stackandqueue

import (
	"container/list"
	"fmt"
	"strconv"
	"testing"
)

/*
https://leetcode.cn/problems/evaluate-reverse-polish-notation/

逆波兰表达式求值
输入：tokens = ["2","1","+","3","*"]
输出：9
解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

解题思路：
如果是运算数，则将其压入栈中。
如果是操作符，则从栈中弹出两个元素。calculateVal = val2.Value.(int) / val1.Value.(int)
注意计算的顺序。
然后将计算的结果再次压入栈中。



*/

func evalRPN(tokens []string) int {
	stack := list.New()
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "+" || token == "-" || token == "*" || token == "/" {
			val1 := stack.Back()
			stack.Remove(val1)
			val2 := stack.Back()
			stack.Remove(val2)

			calculateVal := 0
			switch token {
			case "+":
				calculateVal = val1.Value.(int) + val2.Value.(int)
			case "-":
				calculateVal = val2.Value.(int) - val1.Value.(int)
			case "*":
				calculateVal = val2.Value.(int) * val1.Value.(int)
			case "/":
				calculateVal = val2.Value.(int) / val1.Value.(int)
			}
			stack.PushBack(calculateVal)
		} else {
			num, _ := strconv.Atoi(token)
			stack.PushBack(num)
		}
	}
	return stack.Back().Value.(int)
}

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	res := evalRPN(tokens)
	fmt.Println(res)

	tokens2 := []string{"4", "13", "5", "/", "+"}
	res2 := evalRPN(tokens2)
	fmt.Println(res2)

}
