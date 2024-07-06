package main

import (
	"fmt"
)

func twoSumII(numbers []int, target int) []int {
	var hashmap = make(map[int]int)

	for i, v := range numbers {
		hashmap[v] = i
	}

	for i, v := range numbers {
		if _, ok := hashmap[target-v]; ok && i != hashmap[target-v] {
			return ([]int{i + 1, hashmap[target-v] + 1})
		}
	}

	return []int{0, 0}
}

func main() {
	fmt.Println(twoSumII([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumII([]int{3, 2, 4}, 6))
	fmt.Println(twoSumII([]int{-1, 0}, -1))
}
