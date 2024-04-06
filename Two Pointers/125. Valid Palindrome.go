package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	/*
		在Golang中，不能直接計算字串內的字元，而採用rune遍歷位元組以及下位字串運算。
		rune等相當於int32，長度為4個bytes，常用於unicode或utf-8字符。
		byte等 int8，立即1bytes，常用ascii字元
	*/
	runeSlice := []rune(s)

	left := 0
	right := len(s) - 1

	for left < right {
		if !unicode.IsLetter(runeSlice[left]) && !unicode.IsNumber(runeSlice[left]) {
			left++
		} else if !unicode.IsLetter(runeSlice[right]) && !unicode.IsNumber(runeSlice[right]) {
			right--
		} else if unicode.ToLower(runeSlice[left]) != unicode.ToLower(runeSlice[right]) {
			return false
		} else {
			left++
			right--
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("amanaplanacanalpanama"))
	fmt.Println(isPalindrome("raceacar"))
	fmt.Println(isPalindrome(""))
}
