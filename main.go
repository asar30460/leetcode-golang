package main

import (
	"fmt"
)

/*
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 這個拉鍊是一個規律，每輪都是先往下到底再往上，如此循環
 * 因此只要處理這兩種情況就行
 */
 func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}

    stringBuilder := make([]string, numRows) // 建立列陣列

	i := 0 
	for _, v := range s {
		// 完成循環一輪，重新歸零
		if i == 2 * numRows - 2 {
			i = 0
		}

		if i < numRows {	// 處理往下情況
			stringBuilder[i] += string(v)
		} else {	// 處理往上情況
			stringBuilder[2 * numRows - 2 - i] += string(v)
		}

		i++
	}

	res := ""
	for i := range stringBuilder {
		res += stringBuilder[i]
	}

	return res
 }

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 2))
	fmt.Println(convert("A", 1))
	fmt.Println(convert("AB", 1))
}
