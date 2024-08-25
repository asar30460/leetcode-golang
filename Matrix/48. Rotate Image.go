package main

import (
	"fmt"
)

func rotate(matrix [][]int) [][]int {
	// 1. transpose
	// 2. swap cols

	// transpose
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// swap cols
	left, right := 0, len(matrix[0])-1
	for left < right {
		for i := 0; i < len(matrix); i++ {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
		left++
		right--
	}

	return matrix
}

func main() {
	fmt.Println(rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	fmt.Println(rotate([][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}))
}
