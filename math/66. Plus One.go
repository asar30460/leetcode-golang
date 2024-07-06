package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		// 如果尾數小於9就直接加一回傳
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}

		// 若尾數等於9則進位變成0，這個操作可以一直循環到處理的位數不是9為止
		// 舉例來說7899可以處理到百位數'8'，將它+1再回傳
		digits[i] = 0
	}

	// 阿如果是什麼9999的話，就直接建立一個新陣列，這個陣列大小多一位數設定為，最大位數設'1'
	// 所以9999就會變成10000
	var res = make([]int, n+1)
	res[0] = 1

	return res
}

func main() {
	fmt.Printf("Case 1: %v\n", plusOne([]int{1, 2, 3}))
	fmt.Printf("Case 2: %v\n", plusOne([]int{4, 3, 2, 1}))
	fmt.Printf("Case 3: %v\n", plusOne([]int{9}))
	// 大數運算測試
	fmt.Printf("Case 3: %v\n", plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}
