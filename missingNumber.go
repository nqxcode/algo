package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber(nums []int) int {
	n := len(nums)

	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 0; i <= n; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return -1
}
