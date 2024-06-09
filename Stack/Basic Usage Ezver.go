package main

import (
	"fmt"
)

type EzStack struct {
	data []byte
	top int 
	size int
}

func newStack(size int) *EzStack {
	return &EzStack{
		data : make([]byte, size),
		top: -1,
		size: size,
	}
}

func (e *EzStack) push(val byte) {
	e.top++
	e.data[e.top]=val
}

func (e *EzStack) pop() {
	e.top--
}

func main() {
	var parenthesesStack = newStack(10)
	parenthesesStack.push(byte('a'))
	parenthesesStack.push(byte('b'))
	parenthesesStack.push(byte('c'))
	parenthesesStack.pop()
	fmt.Println(string(parenthesesStack.data[parenthesesStack.top]))
}