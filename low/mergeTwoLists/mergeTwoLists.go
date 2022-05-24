package main

import "leetcode/common"

func main() {
	list1 := common.BuildList([]int{1, 3, 5, 7, 8, 9})
	list2 := common.BuildList([]int{2, 4, 6, 8, 10})
	common.DumpList(list1)
	common.DumpList(list2)

}
