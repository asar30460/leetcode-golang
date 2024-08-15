package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	l, sum := 0, 0
    min := len(nums) + 1

	// 滑動右邊界
    for r := range nums {
        sum += nums[r]

        // 當sum大於目標時我們不斷縮小左邊界
        // 目的是找出在固定右邊界的前提下，最大壓縮這個滑動視窗的極限
        for sum >= target {
            if r - l + 1 < min {
                min = r - l + 1
            }
            sum -= nums[l]

            l++
        }
    }

    if min == len(nums) + 1 {
        return 0
    }

    return min
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen(213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}))
}