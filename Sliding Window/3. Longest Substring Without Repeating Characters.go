package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	res := 0
	left, right := 0, 0

	char := make(map[byte]bool)
	for right < len(s) {
		for char[s[right]] {
			char[s[left]] = false
			left++
		}

		char[s[right]] = true
		diff := right - left + 1
		if diff > res {
			res = diff
		}

		right++
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
