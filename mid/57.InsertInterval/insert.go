package main

import "sort"

func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func min(x ...int) int {
	sort.Ints(x)
	return x[0]
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	i := 0
	ret := make([][]int, 0)
	for i < len(intervals) && newInterval[0] > intervals[i][1] {
		ret = append(ret, intervals[i])
		i++
	}

	for i < len(intervals) && newInterval[1] >= intervals[i][0] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	ret = append(ret, newInterval)

	for i < len(intervals) {
		ret = append(ret, intervals[i])
		i++
	}

	return ret
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	ret := make([][]int, 0)
	start := 0
	end := 0
	foundStart := false
	for i := 0; i < len(intervals); i++ {

		if newInterval[1] >= intervals[i][0] {
			end = i
		}
		if intervals[i][1] >= newInterval[0] && !foundStart {

			start = i
			foundStart = true
			if end < start {
				end = start
			}
		}

	}

	replace := make([][]int, 0)

	//有交集
	if (newInterval[0] >= intervals[start][0] && newInterval[0] <= intervals[start][1]) ||
		(newInterval[1] <= intervals[end][1] && newInterval[1] >= intervals[end][0]) ||
		(newInterval[0] <= intervals[start][0] && newInterval[1] >= intervals[end][1]) {
		replace = append(replace, []int{min(newInterval[0], intervals[start][0]), max(newInterval[1], intervals[end][1])})
	} else {
		//没有交集
		replace = append(replace, newInterval)

		if newInterval[0] < intervals[start][0] {
			ret = append(ret, intervals[:start]...)
			ret = append(ret, newInterval)
			ret = append(ret, intervals[start:]...)
		} else {
			ret = append(ret, intervals...)
			ret = append(ret, newInterval)
		}

		return ret
	}

	for i := 0; i < len(intervals); i++ {
		if i < start || i > end {
			ret = append(ret, intervals[i])
		} else {
			ret = append(ret, replace...)
			i += end - start
		}
	}

	return ret
}
