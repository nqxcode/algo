package main

// https://leetcode.com/problems/top-k-frequent-elements/description/

import (
	"fmt"
	"sort"
)

// aaaabbcccde, 3 - a, c, b
// aaaabbcccde, 2 - a, c

func main() {
	fmt.Println(topK("aaaabbcccde", 3))
	fmt.Println(topK("aaaabbcccde", 2))
}

type Item struct {
	S     rune
	Count int
}

func topK(input string, k int) []string {
	freq := make(map[rune]int)

	for _, s := range input {
		if _, ok := freq[s]; !ok {
			freq[s] = 1
		} else {
			freq[s]++
		}
	}

	items := make([]Item, 0, len(freq))
	for k, v := range freq {
		items = append(items, Item{S: k, Count: v})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})

	items = items[0:k]

	result := make([]string, 0, k)
	for _, item := range items {
		result = append(result, string(item.S))
	}

	return result
}
