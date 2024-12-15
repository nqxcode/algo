package main

// https://leetcode.com/problems/two-sum/

import "fmt"

// Input: nums = [2,11,7,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func main() {
	result := twoSum([]int{2, 1, 7, 8, 10, 20, -11, -1}, 9)
	fmt.Println(result)
}

func twoSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	indexes := make(map[int]int)

	for i, num := range nums {
		value := target - num

		if index, ok := indexes[num]; ok {
			result = append(result, []int{index, i})
		}

		indexes[value] = i
	}

	return result
}
