package main

import "fmt"

// https://leetcode.com/problems/intersection-of-two-arrays/

func main() {
	nums1 := []int{1, 2, 3, 0, 8, 6}
	nums2 := []int{2, 5, 6}

	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {

	nums1Map := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		nums1Map[v] = struct{}{}
	}

	resultMap := make(map[int]struct{})

	for j := 0; j < len(nums2); j++ {
		if _, ok := nums1Map[nums2[j]]; ok {
			resultMap[nums2[j]] = struct{}{}
		}
	}

	result := make([]int, 0, len(resultMap))

	for k, _ := range resultMap {
		result = append(result, k)
	}

	return result
}
