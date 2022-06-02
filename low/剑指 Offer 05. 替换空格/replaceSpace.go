package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	sourceLen := len(s)
	for _, v := range s {
		if v == ' ' {
			sourceLen += 2
		}
	}
	//fmt.Printf("len = %d \n", sourceLen)
	//source := make([]rune, sourceLen, sourceLen)
	source := ""
	for _, v := range s {
		if v != ' ' {
			source += string(v)
		} else {
			source += "%20"
		}
	}

	return source
}

func main() {
	ret := replaceSpace("We are happy.")

	fmt.Printf("%d %s \n", len(ret), ret)
}
