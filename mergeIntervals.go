package main

// https://leetcode.com/problems/merge-intervals/description/

import (
	"fmt"
	"sort"
)

func main() {
	// Пример 1
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("Input:", intervals1)
	fmt.Println("Output:", mergeIntervals(intervals1))

	// Пример 2
	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println("Input:", intervals2)
	fmt.Println("Output:", mergeIntervals(intervals2))
}

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]

		if intervals[i][0] <= last[1] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
