package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	// 為首次出現的字元設定編號
	word_ids_s := make(map[string]int)
	word_ids_t := make(map[string]int)
	
	number_s := 1 // 當新的字元出現時，這個值+1
	number_t := 1

	for i:=0; i<len(s); i++ {
		if word_ids_s[string(s[i])] == 0 {
			word_ids_s[string(s[i])] = number_s
			number_s++
		}

		if word_ids_t[string(t[i])] == 0 {
			word_ids_t[string(t[i])] = number_t
			number_t++
		}

		// {"foo", "bar"}試想當迭帶第三次時，"foo"的'o'已經出現過，而"bar"的'r'還沒有出現過
		// 相對映射的值一個為1一個為0，就可以檢查出節奏有沒有一致
		if word_ids_s[string(s[i])] != word_ids_t[string(t[i])] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}