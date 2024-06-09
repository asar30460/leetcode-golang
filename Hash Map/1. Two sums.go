package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// This is burte force
	var res = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
			}
		}
	}

	return res
}

func twoSumHm(nums []int, target int) []int {
	// This is hash map
	var res = make([]int, 2)
	var hashmap = make(map[int]int)

	for i, v := range nums {
		hashmap[v] = i
	}

	for i, v := range nums {
		if _, ok := hashmap[target-v]; ok && i != hashmap[target-v] {
			res = []int{i, hashmap[target-v]}
			break
		}
	}

	return res
}

func main() {
	fmt.Println(twoSumHm([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumHm([]int{3, 2, 4}, 6))
	fmt.Println(twoSumHm([]int{3, 3}, 6))
	fmt.Println(twoSumHm([]int{-1, -2, -3, -4, -5}, -8))
}
