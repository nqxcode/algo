package main

import "fmt"

// https://leetcode.com/problems/move-zeroes/description/

func main() {

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	zeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			zeroIndex++
		}
	}
}
