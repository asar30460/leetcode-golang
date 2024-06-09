package main

import (
	"fmt"
	"strings"
)

func keyFound(m (map[byte]string), value string, currentKey string) bool {
	for k, v := range m {
		if v == value && string(k) != currentKey {
			// fmt.Printf("found at key %s\n", string(k))
			return true
		}
	}
	return false
}

func wordPattern(s string, t string) bool {
	words := strings.Split(t, " ")

	if len(s) != len(words) {
		return false
	}

	var pattern = make(map[byte]string)	
	
	for i, v := range s {
		// 檢查是否已經設定值，若有的話看本次的值是否與已設定的值相同
		if pattern[byte(v)] != "" && pattern[byte(v)] != words[i] {
			return false
		}
		
		// 若只檢查上面那個判斷的話，在不同key儲存同一個value也會出錯
		// 舉例：("abba", "dog dog dog dog")
		if keyFound(pattern, words[i], string(v)) {
			return false
		}

		pattern[byte(v)] = words[i]
		// fmt.Printf("%v: %s\n", byte(v), words[i])
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
	fmt.Println(wordPattern("aa", "aa aa aa aa"))
}