package main

import "fmt"

// 0, 2, 4, 3
//    5, 6, 4
//    8, 0, 7

func main() {
	fmt.Println(sumArrays([]int{0, 2, 4, 3}, []int{5, 6, 4}))
}

func sumArrays(aNum, bNum []int) []int {
	result := make([]int, max(len(aNum), len(bNum))+1)

	aLen := len(aNum)
	bLen := len(bNum)
	maxLen := max(len(aNum), len(bNum))

	carry := 0
	for i := 0; i < maxLen; i++ {
		aIndex := aLen - i - 1
		a := 0
		if aIndex >= 0 && aIndex < aLen {
			a = aNum[aIndex]
		}

		bIndex := bLen - i - 1
		b := 0
		if bIndex >= 0 && bIndex < bLen {
			b = bNum[bIndex]
		}

		res := a + b + carry
		c := res % 10

		if res > 9 {
			carry = 1
		} else {
			carry = 0
		}

		result[maxLen-i] = c
	}

	return result
}
