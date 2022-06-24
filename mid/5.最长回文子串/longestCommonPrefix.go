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

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbc"))
}
