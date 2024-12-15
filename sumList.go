package main

// https://leetcode.com/problems/add-two-numbers/description/

import "fmt"

type Element struct {
	V    int
	Next *Element
}

// 0, 2, 4, 3
//    5, 6, 4
//    8, 0, 7

func main() {
	num1 := &Element{V: 3, Next: &Element{V: 4, Next: &Element{V: 2, Next: &Element{V: 0}}}}
	num2 := &Element{V: 4, Next: &Element{V: 6, Next: &Element{V: 5}}}

	printNum(sumList(num1, num2))
}

func sumList(aNum, bNum *Element) *Element {
	result := &Element{}
	current := result

	carry := 0
	for {
		if aNum == nil && bNum == nil && carry == 0 {
			break
		}

		a := 0
		if aNum != nil {
			a = aNum.V
		}

		b := 0
		if bNum != nil {
			b = bNum.V
		}

		total := a + b + carry
		c := total % 10

		if total > 9 {
			carry = 1
		} else {
			carry = 0
		}

		current.V = c
		current.Next = &Element{}
		current = current.Next

		if aNum != nil {
			aNum = aNum.Next
		}

		if bNum != nil {
			bNum = bNum.Next
		}
	}

	return result
}

func printNum(num *Element) {
	for num != nil {
		fmt.Printf("%d ", num.V)
		num = num.Next
	}
}
