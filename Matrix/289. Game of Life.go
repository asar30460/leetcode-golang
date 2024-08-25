package main

import (
	"fmt"
)

func gameOfLife(board [][]int) [][]int {
	newBoard := make([][]int, len(board))
	for i := range newBoard {
		newBoard[i] = make([]int, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			lives := liveNeighbors(board, i, j)
			// fmt.Println(lives)
			if board[i][j] == 1 {
				if lives == 2 || lives == 3 {
					newBoard[i][j] = 1
				} else {
					newBoard[i][j] = 0
				}
			} else {
				if lives == 3 {
					newBoard[i][j] = 1
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = newBoard[i][j]
		}
	}

	return board
}

func liveNeighbors(board [][]int, x int, y int) int {
	lives := 0
	for i := max(0, x-1); i < min(x+2, len(board)); i++ {
		for j := max(0, y-1); j < min(y+2, len(board[i])); j++ {
			if i == x && j == y {
				continue
			}
			if board[i][j] == 1 {
				lives++
			}
		}
	}

	return lives
}

func main() {
	fmt.Println(gameOfLife([][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}))

	fmt.Println(gameOfLife([][]int{
		{1, 1},
		{1, 0},
	}))

}
