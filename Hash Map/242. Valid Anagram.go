package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    s_count := make(map[int]int)
    
    for i := 0; i < len(s); i++ {
        s_count[int(s[i])]++
		// fmt.Printf("%d: %d\n", s[i], s_count[int(s[i])])
    }

	for _, v := range t {
		if s_count[int(v)] == 0 {
			return false
		}
		s_count[int(v)]--
	}

    return true
}
func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}