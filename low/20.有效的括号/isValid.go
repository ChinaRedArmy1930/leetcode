package main

import "fmt"

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "(":
			fallthrough
		case "[":
			fallthrough
		case "{":
			stack = append(stack, string(s[i]))

		case ")":
			if len(stack)-1 < 0 {
				return false
			}
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c != "(" {
				return false
			}

		case "]":
			if len(stack)-1 < 0 {
				return false
			}
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c != "[" {
				return false
			}
		case "}":
			if len(stack)-1 < 0 {
				return false
			}
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c != "{" {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true

}

func main() {
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("){"))
}
