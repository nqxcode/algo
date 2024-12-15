package main

import "fmt"

// https://leetcode.com/problems/sort-an-array/description/?envType=problem-list-v2&envId=merge-sort

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
}

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	b := nums[0]

	r := make([]int, 0)
	l := make([]int, 0)

	for _, v := range nums[1:] {
		if v <= b {
			l = append(l, v)
		} else {
			r = append(r, v)
		}
	}

	result := make([]int, 0)

	result = append(result, sortArray(l)...)
	result = append(result, b)
	result = append(result, sortArray(r)...)

	return result
}
