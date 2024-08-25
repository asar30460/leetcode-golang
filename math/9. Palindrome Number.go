package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var numArr []int

	for x != 0 {
		numArr = append(numArr, x%10)
		x /= 10
	}
	// fmt.Println(numArr)

	left, right := 0, len(numArr)-1

	for left < right {
		if numArr[left] != numArr[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Printf("Case 1: %v\n", isPalindrome(121))
	fmt.Printf("Case 2: %v\n", isPalindrome(-121))
	fmt.Printf("Case 3: %v\n", isPalindrome(10))
	fmt.Printf("Case 4: %v\n", isPalindrome(3456))
}
