package main

import (
	"fmt"
)

// 本題用sliding window解，並設定k為搜尋區間
func summaryRanges(nums []int) []string {
	var res []string

	for i := 0; i < len(nums); i++ {
		j := i

		for j+1 < len(nums) && nums[j]+1 == nums[j+1] {
			j++
		}
		if i == j {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		}
		// fmt.Printf("i: %d, j: %d\n", i, j)

		i = j
	}

	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
