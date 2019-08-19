package operatorstr

import (
	"strconv"
	"strings"
)

func solution(s string) int {
	sList := strings.Split(s, " ")
	stack := []int{}

	for _, o := range sList {
		n, err := strconv.Atoi(o)
		if err == nil {
			stack = append(stack, n)
			continue
		}

		if len(stack) <= 1 && (o == "+" || o == "-") {
			return -1
		}

		switch o {
		case "POP":
			stack = stack[:len(stack)-1]
		case "DUP":
			num := stack[len(stack)-1]
			stack = append(stack, num)
		case "+":
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, n1+n2)
		case "-":
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, n1-n2)
		default:
			return -1
		}
	}

	if len(stack) < 1 {
		return -1
	}

	return stack[len(stack)-1]
}
