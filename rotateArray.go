package main

// https://leetcode.com/problems/rotate-array/

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

func rotate(nums []int, k int) {

	for t := 0; t < k; t++ {
		prevIndex := len(nums) - 1
		for i := 0; i < len(nums)-1; i++ {
			p := nums[prevIndex]
			c := nums[i]

			nums[i] = p
			nums[prevIndex] = c
		}
	}
}
