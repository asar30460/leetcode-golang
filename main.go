package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	words := make(map[string]int, 52) // upper case and lower case

	for _, v := range magazine {
		words[string(v)]++
	}

	for _, v := range ransomNote {
		if words[string(v)] == 0 {
			return false
		}

		words[string(v)]--
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
