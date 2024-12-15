package main

// https://leetcode.com/problems/binary-search/

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 5))
}

func binarySearch(nums []int, target int) int {
	lIndex := 0
	rIndex := len(nums) - 1

	for lIndex <= rIndex {
		i := (lIndex + rIndex) / 2

		n := nums[i]

		if n == target {
			return i
		}

		if n < target {
			lIndex = i + 1
		}

		if n > target {
			rIndex = i - 1
		}
	}

	return -1
}
