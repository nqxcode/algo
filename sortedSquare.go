package main

// https://leetcode.com/problems/squares-of-a-sorted-array/description/

import "fmt"

// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1

	var lMax int
	var rMin int

	even := len(nums)%2 == 0 // четное

	if even {
		lMax = len(nums)/2 - 1
		rMin = len(nums)/2 + 1
	} else {
		lMax = len(nums) / 2
		rMin = len(nums) / 2
	}

	result := make([]int, 0, len(nums))

	for {
		if l <= lMax && r >= rMin {
			lSquare := nums[l] * nums[l]
			rSquare := nums[r] * nums[r]

			if l != r {
				if lSquare > rSquare {
					result = append([]int{rSquare, lSquare}, result...)
				} else {
					result = append([]int{lSquare, rSquare}, result...)
				}
			} else {
				result = append([]int{lSquare}, result...)
			}

			l++
			r--
		} else {
			break
		}
	}

	return result
}
