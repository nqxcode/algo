package main

import "fmt"

// https://leetcode.com/problems/merge-sorted-array/

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)

	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}

	i := 0
	j := m

	for {
		if i == m || j == m+n {
			break
		}

		if nums1[i] <= nums1[j] {
			i++
		} else {

			t := nums1[i]
			nums1[i] = nums1[j]
			nums1[j] = t

			j++
		}
	}
}
