package main

import (
	"fmt"
	"strings"
	"time"
)

func checkOk(kuohao string) bool {
	stack := make([]uint8, 0)
	if strings.Count(kuohao, "(") !=
		strings.Count(kuohao, ")") {
		return false
	}

	for i := 0; i < len(kuohao); i++ {
		if kuohao[i] == '(' {
			stack = append(stack, '(')
		}

		if kuohao[i] == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func traversal(index int, allkuohao *[]string, kuohao *string, ret *map[string]int) {
	if index >= len(*allkuohao) {
		if (*kuohao)[len(*kuohao)-1] == '(' {
			return
		}
		if _, ok := (*ret)[*kuohao]; !ok {
			if checkOk(*kuohao) {
				(*ret)[*kuohao] = 1
			}
		}

		return
	}

	kh := []string{"(", ")"}

	for i := 0; i < len(kh); i++ {
		if strings.Count(*kuohao, "(") > (len(*allkuohao)/2) ||
			strings.Count(*kuohao, ")") > (len(*allkuohao)/2) {
			break
		}

		if strings.Count(*kuohao, ")") >
			strings.Count(*kuohao, "(") {
			break
		}

		(*kuohao) += kh[i]
		traversal(index+1, allkuohao, kuohao, ret)
		(*kuohao) = (*kuohao)[:len(*kuohao)-1]
	}

}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	resultMap := make(map[string]int, 0)
	//首先生成括号数组
	allkuohao := make([]string, 0)
	for i := 0; i < n; i++ {
		allkuohao = append(allkuohao, "(")
		allkuohao = append(allkuohao, ")")
	}

	kuohao := "("
	traversal(1, &allkuohao, &kuohao, &resultMap)

	for k, _ := range resultMap {
		result = append(result, k)
	}

	return result
}

func main() {
	now := time.Now()
	for _, v := range generateParenthesis(3) {
		fmt.Println(v)
	}

	fmt.Println(time.Now().Sub(now).Seconds())

}
