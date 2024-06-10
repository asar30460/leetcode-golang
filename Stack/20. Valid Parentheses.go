package main

import (
	"fmt"
)

type EzStack struct {
	data []byte
	top  int
	size int
}

func newStack(size int) *EzStack {
	return &EzStack{
		data: make([]byte, size),
		top:  -1,
		size: size,
	}
}

func (e *EzStack) push(val byte) {
	e.top++
	e.data[e.top] = val
}

func (e *EzStack) pop() {
	e.top--
}

func isValid(s string) bool {
	var parenthesesStack = newStack(len(s))

	for _, v := range s {
		switch v {
		case '{':
			parenthesesStack.push('}')
		case '(':
			parenthesesStack.push(')')
		case '[':
			parenthesesStack.push(']')
		default:
			top := parenthesesStack.top
			// "top == -1" 表示右括號的數量超過左括號，不合理
			if top == -1 {
				return false
			}
			// "byte(v) != parenthesesStack.data[top]" 表示左括號未被正確密合
			if byte(v) != parenthesesStack.data[top] {
				return false
			}

			parenthesesStack.pop()
		}
	}

	return parenthesesStack.top == -1
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("("))
	fmt.Println(isValid("(()"))
	fmt.Println(isValid(")"))
}
