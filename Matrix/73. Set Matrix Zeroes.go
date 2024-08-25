package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) [][]int {
	rows := make(map[int]bool) // 記錄橫排為0的idx
	cols := make(map[int]bool) // 紀錄直排為0的idx
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i, _ := range rows {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = 0
		}
	}

	for i, _ := range cols {
		for j := 0; j < len(matrix); j++ {
			matrix[j][i] = 0
		}
	}

	return matrix
}

func main() {
	fmt.Println(setZeroes([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}))

	fmt.Println(setZeroes([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{3, 4, 5, 2},
	}))
}
