package main

import (
	"fmt"
	"sort"
)

/*
 * 思考一下，3sums == 0 與 2sums == target的關係其實很雷同
 * 因為我們可以將3sums拆解成 2sums + 1sums  -> 2sums = -1sums（a + b = -c）
 * 如此一來我們就可以用 2sums + 1sums 來解決 3sums
 * 不過這題有限制答案不能有重複的組合，也就是說只要不存在重複的target（-c）就行
 * 所以我們可以先對題目給的陣列由小而大進行排序
 * 且設最左邊的數字為a，一旦a出現連續就跳過
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		// 若出現連續相同target，直接跳過本次迭代
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		nestedRes := twoSum(nums[i+1:], -nums[i])
		res = append(res, nestedRes...)
	}

	return res
}

func twoSum(nums []int, target int) [][]int {
	var res [][]int

	var numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}

	for i, v := range nums {
		if j, ok := numMap[target-v]; ok && i < j {
			inner := []int{-target, nums[i], nums[j]}
			if !contains(res, inner) {
				res = append(res, inner)
			}
		}
	}

	return res
}

func contains(slice [][]int, element []int) bool {
	for _, item := range slice {
		if item[0] == element[0] && item[1] == element[1] && item[2] == element[2] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
