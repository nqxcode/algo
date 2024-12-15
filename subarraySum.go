package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}

func subarraySum(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == k {
			count++
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
				break
			}
		}
	}

	return count
}
