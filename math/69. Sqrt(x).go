package main

import (
	"fmt"
)

func mySqrt(x int) int {
    if x <= 1 {
        return x
    }

    var res int
    for i:=0; i<=x; i++ {
        if i*i > x{
            res = i - 1
            break
        }
    }

    return res
}

func main() {
	fmt.Printf("Case 1: %v\n", mySqrt(1))
	fmt.Printf("Case 2: %v\n", mySqrt(2))
	fmt.Printf("Case 3: %v\n", mySqrt(4))
	fmt.Printf("Case 4: %v\n", mySqrt(8))
}
