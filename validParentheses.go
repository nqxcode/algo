package main

import "fmt"

func main() {
	tests := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	for _, test := range tests {
		fmt.Printf("isValid(%q) = %v\n", test, IsValidParentheses(test))
	}
}

func IsValidParentheses(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if opening, ok := brackets[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
