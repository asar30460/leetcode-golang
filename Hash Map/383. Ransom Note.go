package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	words := make(map[string]int, 52) // upper case and lower case

	for _, v := range magazine {
		words[string(v)]++
	}

	/*
	只要檢測ransomNote字元出現字數超過magazine，表示ransomNote不能構成magazine。
	即便magazine有的字元在ransomNote沒有出現也無所謂
	*/
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
