package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		// 若出現連續相同target，直接跳過本次迭代
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		/* 以下判斷就剩下「2sums == -nums[i]?...等式(一)」，用左右指標往內夾搜尋
		 * 而nums是排序過的，2sums < -nums[i]時，增加left，同理，2sums > -nums[i]時，減少right
		 * 此外我們一樣不希望2sums也出現重複組合
		 * 所以我們要確認在左右指標有更動過（非第一次迭帶），判斷是否與上一次的結果重複，如是，就跳過本次迭代
		*/
		left, right := i+1, len(nums)-1
		for left < right {
			if left != i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			if right != len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			if nums[left]+nums[right] == -nums[i] {
				res = append(res, []int{nums[i], nums[left], nums[right]})
			}

			if nums[left]+nums[right] < -nums[i] {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
