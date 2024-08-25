package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1 // 設定上下左右邊界

	// 任一邊走到底時發現邊界緊密，結束迭代
	for {
		// 往右
		fmt.Print(">: ")
		for i := left; i <= right; i++ {
			fmt.Printf("%v, ", matrix[top][i])
			res = append(res, matrix[top][i])
		}
		top++
		fmt.Println()
		if top > bottom {
			break
		}

		// 往下
		fmt.Print("v: ")
		for i := top; i <= bottom; i++ {
			fmt.Printf("%v, ", matrix[i][right])
			res = append(res, matrix[i][right])
		}
		right--
		fmt.Println()
		if left > right {
			break
		}

		// 往左
		fmt.Print("<: ")
		for i := right; i >= left; i-- {
			fmt.Printf("%v, ", matrix[bottom][i])
			res = append(res, matrix[bottom][i])
		}
		bottom--
		fmt.Println()
		if top > bottom {
			break
		}

		// 往上
		fmt.Print("^: ")
		for i := bottom; i >= top; i-- {
			fmt.Printf("%v, ", matrix[i][left])
			res = append(res, matrix[i][left])
		}
		left++
		fmt.Println()
		if left > right {
			break
		}
	}

	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}
