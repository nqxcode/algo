package main

// https://leetcode.com/problems/monotonic-array/description/

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(isMonotonic([]int{1, 3, 2}))
}

func isMonotonic(nums []int) bool {
	isAsc := -1
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}

		if nums[i] <= nums[i+1] {
			isAsc = 1
			continue
		} else {
			isAsc = 0
			break
		}
	}

	isDesc := -1
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}

		if nums[i] >= nums[i+1] {
			isDesc = 1
			continue
		} else {
			isDesc = 0
			break
		}
	}

	return isAsc == 1 || isDesc == 1
}
