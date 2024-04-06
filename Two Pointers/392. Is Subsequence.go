package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	idx_s := 0
	idx_t := 0
    
	for idx_s < len(s) && idx_t < len(t) {
		if s[idx_s] == t[idx_t] {
			idx_s++
			idx_t++
		} else {
			idx_t++
		}
	}

	return idx_s == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("b", "abc"))
	fmt.Println(isSubsequence("aaaaaa", "bbaaaa"))
}
