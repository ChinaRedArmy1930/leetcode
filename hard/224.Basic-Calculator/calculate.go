package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const N = 100010

type stack struct {
	bottom int
	top    int
	vals   [N]interface{}
}

func constructor() stack {
	return stack{
		bottom: 0,
		top:    0,
		vals:   [N]interface{}{},
	}
}

func (s *stack) pop() {
	s.top--
}

func (s *stack) query() interface{} {
	return s.vals[s.top]
}

func (s *stack) push(x interface{}) {
	s.top++
	s.vals[s.top] = x
}

func (s *stack) empty() bool {
	return s.top <= s.bottom
}

func eval(nums, op *stack) {
	b := nums.query().(int)
	nums.pop()
	a := nums.query().(int)
	nums.pop()
	var x int

	switch op.query().(rune) {
	case '+':
		x = a + b
	case '-':
		x = a - b
	case '*':
		x = a * b
	case '/':
		x = a / b
	}
	op.pop()

	nums.push(x)
}

func calculate(s string) int {
	result := 0

	pr := make(map[rune]int, 0)
	pr['+'] = 1
	pr['-'] = 1
	pr['*'] = 2
	pr['/'] = 2

	nums := constructor()
	op := constructor()
	//第一个数字为负数，则在前补一个0
	if s[0] == '-' {
		nums.push(0)
	}

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			j := i
			x := 0
			for j < len(s) && unicode.IsDigit(rune(s[j])) {
				x = x*10 + int(s[j]-'0')
				j++
			}
			i = j - 1
			nums.push(x)
		} else if s[i] == '(' {
			op.push(rune(s[i]))
			if s[i+1] == '-' {
				nums.push(0)
			}

		} else if s[i] == ')' {
			for op.query().(rune) != '(' {
				eval(&nums, &op)
			}
			op.pop()
		} else if s[i] == ' ' {
			continue
		} else {
			for !op.empty() && op.query().(rune) != '(' && pr[op.query().(rune)] >= pr[rune(s[i])] {
				eval(&nums, &op)
			}
			op.push(rune(s[i]))
		}
	}

	for !op.empty() {
		eval(&nums, &op)
	}
	result = nums.query().(int)

	return result
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(reader, &s)

	fmt.Println(s)
	fmt.Println(calculate(s))
}
