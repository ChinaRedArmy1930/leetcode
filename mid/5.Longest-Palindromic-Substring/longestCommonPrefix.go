package main

import (
	"fmt"
)

func prefix(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	//return s[l+1 : l+1+(r-(l+1))]
	return s[l+1 : r]
}

func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		t1 := prefix(s, i, i)
		t2 := prefix(s, i, i+1)
		tmp := ""
		if len(t1) > len(t2) {
			tmp = t1
		} else {
			tmp = t2
		}

		if len(max) < len(tmp) {
			max = tmp
		}
	}

	return max
}

func dump(nums [][]int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func longestPalindromeDp(s string) string {
	//dp[i][j] 代表s[i:j] 为回文串
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		t := make([]int, len(s))
		dp[i] = t
	}

	//对于i == j 则代表只有一个字符 肯定是回文串
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == j {
				dp[i][j] = 1
			}
		}
	}

	left := 0
	right := 0
	max := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				continue
			}

			if j-i <= 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				} else {
					//此处已经计算出对应的值了, 不应该在走获取左下数值的操作
					continue
				}
			}

			if dp[i+1][j-1] == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			}

			//这里更新最大值
			if dp[i][j] == 1 && j-i+1 > max {
				max = j - i + 1
				left = i
				right = j
			}
		}
	}

	for _, v := range dp {
		fmt.Println(v)
	}

	fmt.Println()

	return s[left : right+1]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbc"))
	fmt.Println(longestPalindromeDp("babad"))
	fmt.Println(longestPalindromeDp("cbbc"))
}
